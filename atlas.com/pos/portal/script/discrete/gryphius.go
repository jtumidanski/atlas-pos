package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Gryphius struct {
}

func (p Gryphius) Name() string {
	return "gryphius"
}

func (p Gryphius) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(240020101, "out00")
	return true
}
