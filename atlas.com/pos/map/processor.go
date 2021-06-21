package _map

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

func MonsterCount(l logrus.FieldLogger, c script.Context) func(mapId uint32) int {
	return func(mapId uint32) int {
		return 0
	}
}
