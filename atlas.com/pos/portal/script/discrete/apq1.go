package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq1 struct {
}

func (p Apq1) Name() string {
	return "apq1"
}

func (p Apq1) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010400, 0)
	return true
}
