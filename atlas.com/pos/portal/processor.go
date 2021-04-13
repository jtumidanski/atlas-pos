package portal

import (
	"atlas-pos/domain"
	"atlas-pos/kafka/producers"
	"atlas-pos/portal/blocked"
	"atlas-pos/portal/script"
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"context"
	"log"
	"strconv"
)

type Processor struct {
	l *log.Logger
}

func NewProcessor(l *log.Logger) *Processor {
	return &Processor{l}
}

func (p *Processor) EnterPortal(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	portal, err := p.getMapPortalById(mapId, portalId)
	if err != nil {
		p.l.Printf("[WARN] unable to locate portal %d for map %d.", portalId, mapId)
		return
	}
	p.enterPortal(worldId, channelId, characterId, mapId, portal)
}

func (p *Processor) getMapPortalById(mapId uint32, portalId uint32) (*domain.PortalModel, error) {
	resp, err := requests.MapInformation().GetPortalById(mapId, portalId)
	if err != nil {
		return nil, err
	}

	d := resp.Data()
	aid, err := strconv.ParseUint(d.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	a := makePortal(uint32(aid), d.Attributes)
	return &a, nil
}

func (p *Processor) getMapPortalByName(mapId uint32, portalName string) (*domain.PortalModel, error) {
	resp, err := requests.MapInformation().GetPortalByName(mapId, portalName)
	if err != nil {
		return nil, err
	}

	d := resp.Data()
	aid, err := strconv.ParseUint(d.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	a := makePortal(uint32(aid), d.Attributes)
	return &a, nil
}

func (p *Processor) enterPortal(worldId byte, channelId byte, characterId uint32, mapId uint32, portal *domain.PortalModel) {
	// TODO check portal delay
	if portal == nil || blocked.GetCache().Blocked(characterId, portal.ScriptName()) {
		producers.EnableActions(p.l, context.Background()).Emit(characterId)
		return
	}

	changed := false
	if portal.ScriptName() != "" {
		// execute portal script
		s, err := script.GetScriptRegistry().GetScript(portal.ScriptName())
		if err != nil {
			p.l.Printf("[WARN] missing script %s for portal %d.", portal.ScriptName(), portal.Id())
			producers.EnableActions(p.l, context.Background()).Emit(characterId)
			return
		}

		c := script.NewContext(worldId, channelId, characterId, mapId, portal)

		if ps, ok := (*s).AsPortalScript(p.l, c); ok {
			changed = ps.Enter()
		} else {
			p.l.Printf("[WARN] script retrieved named %s was not a PortalScript.", portal.ScriptName())
		}
	} else if portal.TargetMapId() != 999999999 {
		// invalidate map change if trying to move with chalkboard open, and target is a free market map.

		toPortal, err := p.getMapPortalByName(portal.TargetMapId(), portal.Target())
		if err != nil {
			p.l.Printf("[INFO] unable to locate portal target %s for map %d, defaulting to portal 0.", portal.Target(), portal.TargetMapId())
			toPortal, err = p.getMapPortalById(portal.TargetMapId(), 0)
			if err != nil {
				p.l.Printf("[ERROR] unable to locate portal 0 for map %d, is the destination map invalid?", portal.TargetMapId())
				producers.EnableActions(p.l, context.Background()).Emit(characterId)
				return
			}
		}

		producers.ChangeMap(p.l, context.Background()).Emit(worldId, channelId, characterId, portal.TargetMapId(), toPortal.Id())
		changed = true
	}

	if !changed {
		producers.EnableActions(p.l, context.Background()).Emit(characterId)
	}
}

func makePortal(id uint32, attr attributes.PortalAttributes) domain.PortalModel {
	return domain.NewPortalModel(id, attr.Name, attr.Target, attr.TargetMap, attr.Type, attr.X, attr.Y, attr.ScriptName)
}
