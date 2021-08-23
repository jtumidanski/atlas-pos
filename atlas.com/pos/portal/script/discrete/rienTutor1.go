package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RienTutor1 struct {
}

func (p RienTutor1) Name() string {
	return "rienTutor1"
}

func (p RienTutor1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestCompleted(l, c)(21010) {
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(140090200, 1)
	return true
}