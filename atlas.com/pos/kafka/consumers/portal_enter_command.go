package consumers

import (
	"atlas-pos/kafka/handler"
	"atlas-pos/kafka/producers"
	"atlas-pos/portal"
	"atlas-pos/portal/blocked"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/registry"
	"github.com/sirupsen/logrus"
)

type PortalEnterCommand struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func PortalEnterCommandCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &PortalEnterCommand{}
	}
}

func HandlePortalEnterCommand() handler.EventHandler {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*PortalEnterCommand); ok {
			model, err := portal.GetById(event.MapId, event.PortalId)
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

func enterPortal(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *portal.Model) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *portal.Model) {
		// TODO check portal delay
		if model == nil || blocked.GetCache().Blocked(characterId, model.ScriptName()) {
			producers.EnableActions(l)(characterId)
			return
		}

		changed := false
		if model.ScriptName() != "" {
			// execute portal script
			s, err := registry.GetScriptRegistry().GetScript(model.ScriptName())
			if err != nil {
				l.WithError(err).Warnf("Missing script %s for portal %d.", model.ScriptName(), model.Id())
				producers.EnableActions(l)(characterId)
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

			toPortal, err := portal.GetByName(model.TargetMapId(), model.Target())
			if err != nil {
				l.WithError(err).Infof("Unable to locate portal target %s for map %d, defaulting to portal 0.", model.Target(), model.TargetMapId())
				toPortal, err = portal.GetById(model.TargetMapId(), 0)
				if err != nil {
					l.WithError(err).Errorf("Unable to locate portal 0 for map %d, is the destination map invalid?", model.TargetMapId())
					producers.EnableActions(l)(characterId)
					return
				}
			}

			producers.ChangeMap(l)(worldId, channelId, characterId, model.TargetMapId(), toPortal.Id())
			changed = true
		}

		if !changed {
			producers.EnableActions(l)(characterId)
		}
	}
}
