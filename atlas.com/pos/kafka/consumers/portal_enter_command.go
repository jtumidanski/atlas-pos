package consumers

import (
	"atlas-pos/domain"
	"atlas-pos/kafka/producers"
	"atlas-pos/portal"
	"atlas-pos/portal/blocked"
	"atlas-pos/portal/script"
	"context"
	"github.com/sirupsen/logrus"
)

type PortalEnterCommand struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func PortalEnterCommandCreator() EmptyEventCreator {
	return func() interface{} {
		return &PortalEnterCommand{}
	}
}

func HandlePortalEnterCommand() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*PortalEnterCommand); ok {
			p := portal.NewProcessor(l)

			model, err := p.GetMapPortalById(event.MapId, event.PortalId)
			if err != nil {
				l.WithError(err).Warnf("Unable to locate portal %d for map %d.", event.MapId, event.PortalId)
				return
			}
			enterPortal(l)(event.WorldId, event.ChannelId, event.CharacterId, event.MapId, model)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}

func enterPortal(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *domain.PortalModel) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *domain.PortalModel) {
		p := portal.NewProcessor(l)

		// TODO check portal delay
		if model == nil || blocked.GetCache().Blocked(characterId, model.ScriptName()) {
			producers.EnableActions(l, context.Background()).Emit(characterId)
			return
		}

		changed := false
		if model.ScriptName() != "" {
			// execute portal script
			s, err := script.GetScriptRegistry().GetScript(model.ScriptName())
			if err != nil {
				l.WithError(err).Warnf("Missing script %s for portal %d.", model.ScriptName(), model.Id())
				producers.EnableActions(l, context.Background()).Emit(characterId)
				return
			}

			c := script.NewContext(worldId, channelId, characterId, mapId, model)

			if ps, ok := (*s).AsPortalScript(l, c); ok {
				changed = ps.Enter()
			} else {
				l.Warnf("Script retrieved named %s was not a PortalScript.", model.ScriptName())
			}
		} else if model.TargetMapId() != 999999999 {
			// invalidate map change if trying to move with chalkboard open, and target is a free market map.

			toPortal, err := p.GetMapPortalByName(model.TargetMapId(), model.Target())
			if err != nil {
				l.WithError(err).Infof("Unable to locate portal target %s for map %d, defaulting to portal 0.", model.Target(), model.TargetMapId())
				toPortal, err = p.GetMapPortalById(model.TargetMapId(), 0)
				if err != nil {
					l.WithError(err).Errorf("Unable to locate portal 0 for map %d, is the destination map invalid?", model.TargetMapId())
					producers.EnableActions(l, context.Background()).Emit(characterId)
					return
				}
			}

			producers.ChangeMap(l, context.Background()).Emit(worldId, channelId, characterId, model.TargetMapId(), toPortal.Id())
			changed = true
		}

		if !changed {
			producers.EnableActions(l, context.Background()).Emit(characterId)
		}
	}
}
