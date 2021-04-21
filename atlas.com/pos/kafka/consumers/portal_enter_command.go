package consumers

import (
	"atlas-pos/portal"
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
			portal.NewProcessor(l).EnterPortal(event.WorldId, event.ChannelId, event.CharacterId, event.MapId, event.PortalId)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
