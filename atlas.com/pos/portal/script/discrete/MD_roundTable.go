package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDRoundTable struct {
}

func (p MDRoundTable) Name() string {
	return "MD_roundTable"
}

func (p MDRoundTable) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(240020500)
	dungeonId := uint32(240020512)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
