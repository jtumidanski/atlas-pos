package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq3 struct {
}

func (p Apq3) Name() string {
	return "apq3"
}

func (p Apq3) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010600, 0)
	return true
}
