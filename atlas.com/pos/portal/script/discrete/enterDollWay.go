package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterDollWay struct {
}

func (p EnterDollWay) Name() string {
	return "enterDollWay"
}

func (p EnterDollWay) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(20730) || script.QuestCompleted(l, c)(21734) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(105070300, 3)
		return true
	}
	if script.QuestStarted(l, c)(21734) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(910510100, 0)
		return true
	}
	script.SendPinkNotice(l, c)("OMINOUS_POWER")
	return false
}
