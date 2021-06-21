package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorOut3 struct {
}

func (p AranTutorOut3) Name() string {
	return "aranTutorOut3"
}

func (p AranTutorOut3) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.TeachSkill(l, c)(20000016, 0, -1, -1)
	character.TeachSkill(l, c)(20000016, 1, 0, -1)
	character.PlayPortalSound(l)
	character.WarpById(l, c)(914000220, 1)
	return true
}
