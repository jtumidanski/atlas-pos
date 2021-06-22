package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaU61 struct {
}

func (p MasteriaU61) Name() string {
	return "Masteria_U6_1"
}

func (p MasteriaU61) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610010002, "sU3_1")
		return false
	}
	return true
}
