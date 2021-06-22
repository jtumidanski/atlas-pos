package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDMushroom struct {
}

func (p MDMushroom) Name() string {
	return "MD_mushroom"
}

func (p MDMushroom) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(105050100)
	dungeonId := uint32(105050101)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
