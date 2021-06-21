package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorOut1 struct {
}

func (p AranTutorOut1) Name() string {
	return "aranTutorOut1"
}

func (p AranTutorOut1) Enter(l logrus.FieldLogger, c script.Context) bool {

	if script.QuestStarted(l, c)(21000) {
		script.TeachSkill(l, c)(20000017, 0, -1, -1)
		script.TeachSkill(l, c)(20000018, 0, -1, -1)
		script.TeachSkill(l, c)(20000018, 1, 0, -1)
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(914000200, 1)
		return true
	} else {
		script.SendPinkNotice(l, c)("ARAN_TUTORIAL_EXIT")
		return false
	}
}
