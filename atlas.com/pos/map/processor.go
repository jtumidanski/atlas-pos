package _map

import (
	"atlas-pos/monster"
	"github.com/sirupsen/logrus"
)

func DojoPoints(mapId uint32) uint32 {
	points := 1 + ((mapId-1)/100%100)/6
	if !DojoPartyArea(mapId) {
		points++
	}
	return points
}

func DojoPartyArea(mapId uint32) bool {
	return mapId >= 925030100 && mapId < 925040000
}

func MonsterCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) int {
	return func(worldId byte, channelId byte, mapId uint32) int {
		return 0
	}
}

func MonsterById(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32) *monster.Model {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32) *monster.Model {
		return nil
	}
}

func CharactersInMap(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) []uint32 {
	return func(worldId byte, channelId byte, mapId uint32) []uint32 {
		return []uint32{}
	}
}
