package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Apq3 struct {
}

func (p Apq3) Name() string {
	return "apq3"
}

func (p Apq3) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(670010600, 0)
	return true
}
