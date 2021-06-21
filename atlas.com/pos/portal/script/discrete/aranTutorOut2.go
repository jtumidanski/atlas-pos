package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorOut2 struct {
}

func (p AranTutorOut2) Name() string {
	return "aranTutorOut2"
}

func (p AranTutorOut2) Enter(l logrus.FieldLogger, c script.Context) bool {

	character.TeachSkill(l, c)(20000014, 0, -1, -1)
	character.TeachSkill(l, c)(20000015, 0, -1, -1)
	character.TeachSkill(l, c)(20000014, 1, 0, -1)
	character.TeachSkill(l, c)(20000015, 1, 0, -1)
	character.PlayPortalSound(l)
	character.WarpById(l, c)(914000210, 1)
	return true
}
