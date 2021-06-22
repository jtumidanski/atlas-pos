package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type MDDrakeRoom struct {
}

func (p MDDrakeRoom) Name() string {
	return "MD_drakeroom"
}

func (p MDDrakeRoom) Enter(l logrus.FieldLogger, c script.Context) bool {
	baseId := uint32(105090311)
	dungeonId := uint32(105090320)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, c)(baseId, dungeonId, dungeons)
}
