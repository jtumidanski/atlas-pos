package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AriantMount struct {
}

func (p AriantMount) Name() string {
	return "ariantMount"
}

func (p AriantMount) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(980010020, 0)
	return true
}
