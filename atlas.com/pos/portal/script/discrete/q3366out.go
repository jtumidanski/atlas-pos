package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Q3366Out struct {
}

func (p Q3366Out) Name() string {
	return "q3366out"
}

func (p Q3366Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(926130100, "in00")
	return true
}
