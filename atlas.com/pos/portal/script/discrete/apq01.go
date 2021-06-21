package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq01 struct {
}

func (p Apq01) Name() string {
	return "apq01"
}

func (p Apq01) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010301, 0)
	return true
}
