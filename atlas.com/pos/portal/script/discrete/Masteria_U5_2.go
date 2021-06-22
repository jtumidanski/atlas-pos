package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaU52 struct {
}

func (p MasteriaU52) Name() string {
	return "Masteria_U5_2"
}

func (p MasteriaU52) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610010004, "U5_1")
		return false
	}
	return true
}
