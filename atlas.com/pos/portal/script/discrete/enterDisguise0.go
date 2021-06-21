package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterDisguise0 struct {
}

func (p EnterDisguise0) Name() string {
	return "enterDisguise0"
}

func (p EnterDisguise0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(20301) ||
		script.QuestStarted(l, c)(20302) ||
		script.QuestStarted(l, c)(20303) ||
		script.QuestStarted(l, c)(20304) ||
		script.QuestStarted(l, c)(20305) {
		if script.HasItem(l, c)(4032179) {
			script.PlayPortalSound(l, c)
			script.WarpByName(l, c)(130010000, "east00")
			return true
		}
		script.SendPinkNotice(l, c)("NEED_PERMIT")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(130010000, "east00")
	return true
}
