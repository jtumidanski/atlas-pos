package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutNix2 struct {
}

func (p OutNix2) Name() string {
	return "outNix2"
}

func (p OutNix2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020401,"in00")
	return true
}
