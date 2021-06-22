package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Mayong struct {
}

func (p Mayong) Name() string {
	return "mayong"
}

func (p Mayong) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020401, "out00")
	return true
}
