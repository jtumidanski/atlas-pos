package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDProtect struct {
}

func (p MDProtect) Name() string {
	return "MD_protect"
}

func (p MDProtect) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(240040520)
	dungeonId := uint32(240040900)
	dungeons := uint8(19)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
