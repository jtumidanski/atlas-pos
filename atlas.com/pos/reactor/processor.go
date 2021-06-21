package reactor

import "github.com/sirupsen/logrus"

func ByName(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorName string) *Model {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) *Model {
		return nil
	}
}
