package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Q3368Out struct {
}

func (p Q3368Out) Name() string {
	return "q3368out"
}

func (p Q3368Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(926130100, "in02")
	return true
}
