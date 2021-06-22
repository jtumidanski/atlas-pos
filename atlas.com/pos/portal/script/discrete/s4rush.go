package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4Rush struct {
}

func (p S4Rush) Name() string {
	return "s4rush"
}

func (p S4Rush) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(6110) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(910500100, 0)
		return true
	}
	script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
