package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1D struct {
}

func (p MasteriaCM1D) Name() string {
	return "Masteria_CM1_D"
}

func (p MasteriaCM1D) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610020000, "CM1_E")
		return false
	}
	return true
}
