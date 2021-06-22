package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InNix2 struct {
}

func (p InNix2) Name() string {
	return "inNix2"
}

func (p InNix2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020600,"out01")
	return true
}
