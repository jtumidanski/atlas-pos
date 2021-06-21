package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq2 struct {
}

func (p Apq2) Name() string {
	return "apq2"
}

func (p Apq2) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010500, 0)
	return true
}
