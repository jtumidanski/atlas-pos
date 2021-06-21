package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq00 struct {
}

func (p Apq00) Name() string {
	return "apq00"
}

func (p Apq00) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010300, 0)
	return true
}
