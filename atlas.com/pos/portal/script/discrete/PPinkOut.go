package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type PPinkOut struct {
}

func (p PPinkOut) Name() string {
	return "PPinkOut"
}

func (p PPinkOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(270050000)
	return true
}
