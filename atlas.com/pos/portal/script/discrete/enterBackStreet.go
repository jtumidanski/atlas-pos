package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterBackStreet struct {
}

func (p EnterBackStreet) Name() string {
	return "enterBackStreet"
}

func (p EnterBackStreet) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestActive(l, c)(21747) || script.QuestActive(l, c)(21744) && script.QuestCompleted(l, c)(21745) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(925040000, 0)
		return true
	}
	script.SendPinkNotice(l, c)("NO_PERMISSION_TO_ENTER")
	return false
}
