package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorAloneX struct {
}

func (p AranTutorAloneX) Name() string {
	return "aranTutorAloneX"
}

func (p AranTutorAloneX) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(914000100, 1)
	return true
}
