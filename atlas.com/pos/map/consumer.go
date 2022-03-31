package _map

import (
	"atlas-pos/character"
	"atlas-pos/kafka"
	"atlas-pos/portal"
	"atlas-pos/portal/blocked"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerNameEnterPortal = "enter_portal_command"
	topicTokenEnterPortal   = "TOPIC_ENTER_PORTAL"
)

func EnterPortalConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[enterPortalCommand](consumerNameEnterPortal, topicTokenEnterPortal, groupId, handleEnterPortal())
}

type enterPortalCommand struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
	CharacterId uint32 `json:"characterId"`
}

func handleEnterPortal() kafka.HandlerFunc[enterPortalCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command enterPortalCommand) {
		model, err := portal.ByIdModelProvider(l, span)(command.MapId, command.PortalId)()
		if err != nil {
			l.WithError(err).Warnf("Unable to locate portal %d for map %d.", command.MapId, command.PortalId)
			return
		}
		enterPortal(l, span)(command.WorldId, command.ChannelId, command.CharacterId, command.MapId, model)
	}
}

func enterPortal(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *portal.Model) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, model *portal.Model) {
		// TODO check portal delay
		if model == nil || blocked.GetCache().Blocked(characterId, model.ScriptName()) {
			character.EnableActions(l, span)(worldId, channelId, characterId)
			return
		}

		changed := false
		if model.ScriptName() != "" {
			// execute portal script
			s, err := script.GetRegistry().GetScript(model.ScriptName())
			if err != nil {
				l.WithError(err).Warnf("Missing script %s for portal %d.", model.ScriptName(), model.Id())
				character.EnableActions(l, span)(worldId, channelId, characterId)
				return
			}

			c := script.NewContext(worldId, channelId, characterId, mapId, model)

			if ps, ok := (*s).AsPortalScript(l, span, c); ok {
				changed = ps.Enter()
			} else {
				l.Warnf("Script retrieved named %s was not a PortalScript.", model.ScriptName())
			}
		} else if model.TargetMapId() != 999999999 {
			// invalidate map change if trying to move with chalkboard open, and target is a free market map.

			toPortal, err := portal.ByNameModelProvider(l, span)(model.TargetMapId(), model.Target())()
			if err != nil {
				l.WithError(err).Infof("Unable to locate portal target %s for map %d, defaulting to portal 0.", model.Target(), model.TargetMapId())
				toPortal, err = portal.ByIdModelProvider(l, span)(model.TargetMapId(), 0)()
				if err != nil {
					l.WithError(err).Errorf("Unable to locate portal 0 for map %d, is the destination map invalid?", model.TargetMapId())
					character.EnableActions(l, span)(worldId, channelId, characterId)
					return
				}
			}

			portal.WarpById(l, span)(worldId, channelId, characterId, model.TargetMapId(), toPortal.Id())
			changed = true
		}

		if !changed {
			character.EnableActions(l, span)(worldId, channelId, characterId)
		}
	}
}
