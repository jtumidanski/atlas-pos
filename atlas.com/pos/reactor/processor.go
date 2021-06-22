package reactor

import "github.com/sirupsen/logrus"

func ByName(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorName string) *Model {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) *Model {
		return nil
	}
}

func ById(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorId uint32) *Model {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32) *Model {
		return nil
	}
}

func ForceHit(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, readtorId uint32, state byte) {
	return func(worldId byte, channelId byte, mapId uint32, readtorId uint32, state byte) {

	}
}
