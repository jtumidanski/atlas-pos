package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Pinkin struct {
}

func (p Pinkin) Name() string {
	return "Pinkin"
}

func (p Pinkin) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(270050100)
	return true
}
