package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDRemember struct {
}

func (p MDRemember) Name() string {
	return "MD_remember"
}

func (p MDRemember) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(240040511)
	dungeonId := uint32(240040800)
	dungeons := uint8(19)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
