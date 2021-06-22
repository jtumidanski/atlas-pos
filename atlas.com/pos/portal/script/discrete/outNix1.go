package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutNix1 struct {
}

func (p OutNix1) Name() string {
	return "outNix1"
}

func (p OutNix1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020101,"in00")
	return true
}
