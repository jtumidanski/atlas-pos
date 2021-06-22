package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaB21 struct {
}

func (p MasteriaB21) Name() string {
	return "Masteria_B2_1"
}

func (p MasteriaB21) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610010005, "sU6_1")
		return true
	}
	return false
}
