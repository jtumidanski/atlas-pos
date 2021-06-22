package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDGolem struct {
}

func (p MDGolem) Name() string {
	return "MDlem"
}

func (p MDGolem) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(105040304)
	dungeonId := uint32(105040320)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
