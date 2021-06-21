package discrete

import (
	"atlas-pos/character"
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
		character.TeachSkill(l, c)(20000017, 0, -1, -1)
		character.TeachSkill(l, c)(20000018, 0, -1, -1)
		character.TeachSkill(l, c)(20000018, 1, 0, -1)
		character.PlayPortalSound(l)
		character.WarpById(l, c)(914000200, 1)
		return true
	} else {
		character.SendPinkNotice(l, c)("ARAN_TUTORIAL_EXIT")
		return false
	}
}
