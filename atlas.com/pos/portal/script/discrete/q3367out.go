package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Q3367Out struct {
}

func (p Q3367Out) Name() string {
	return "q3367out"
}

func (p Q3367Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(926130100, "in01")
	return true
}
