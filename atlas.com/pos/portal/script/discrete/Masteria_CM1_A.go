package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1A struct {
}

func (p MasteriaCM1A) Name() string {
	return "Masteria_CM1_A"
}

func (p MasteriaCM1A) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610020000, "CM1_B")
		return false
	}
	return true
}
