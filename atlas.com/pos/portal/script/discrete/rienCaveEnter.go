package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RienCaveEnter struct {
}

func (p RienCaveEnter) Name() string {
	return "rienCaveEnter"
}

func (p RienCaveEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(21201) || script.QuestStarted(l, c)(21302) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(140030000, 1)
		return true
	}
	script.SendPinkNotice(l, c)("SOMETHING_BLOCKING_PORTAL")
	return false
}
