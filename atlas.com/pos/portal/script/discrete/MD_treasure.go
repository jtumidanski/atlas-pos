package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDTreasure struct {
}

func (p MDTreasure) Name() string {
	return "MD_treasure"
}

func (p MDTreasure) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(251010402)
	dungeonId := uint32(251010410)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
