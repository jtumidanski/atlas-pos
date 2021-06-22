package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TempleEnter struct {
}

func (p TempleEnter) Name() string {
	return "templeenter"
}

func (p TempleEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.CancelItem(l, c)(2210016)
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(270000100, "out00")
	return true
}
