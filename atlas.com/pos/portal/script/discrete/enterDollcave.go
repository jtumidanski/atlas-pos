package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterDollcave struct {
}

func (p EnterDollcave) Name() string {
	return "enterDollcave"
}

func (p EnterDollcave) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(20730) || script.QuestCompleted(l, c)(21734) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(105040201, 2)
		return true
	}
	script.OpenNPC(l, c)(1063011, "PupeteerPassword")
	return false
}
