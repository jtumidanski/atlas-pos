package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq02 struct {
}

func (p Apq02) Name() string {
	return "apq02"
}

func (p Apq02) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(670010302, 0)
	return true
}
