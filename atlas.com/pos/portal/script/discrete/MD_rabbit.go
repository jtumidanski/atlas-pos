package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDRabbit struct {
}

func (p MDRabbit) Name() string {
	return "MD_rabbit"
}

func (p MDRabbit) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(221023400)
	dungeonId := uint32(221023401)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
