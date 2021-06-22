package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InNix1 struct {
}

func (p InNix1) Name() string {
	return "inNix1"
}

func (p InNix1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020600,"out00")
	return true
}
