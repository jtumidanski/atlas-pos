package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type OutTemple struct {
}

func (p OutTemple) Name() string {
	return "outTemple"
}

func (p OutTemple) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.UseItem(l, c)(2210016)
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(200090510, 0)
	return true
}
