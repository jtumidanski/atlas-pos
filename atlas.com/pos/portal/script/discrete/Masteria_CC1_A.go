package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCC1A struct {
}

func (p MasteriaCC1A) Name() string {
	return "Masteria_CC1_A"
}

func (p MasteriaCC1A) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(610020015, "CC6_A")
	return true
}
