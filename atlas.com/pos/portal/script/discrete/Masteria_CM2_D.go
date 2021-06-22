package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2D struct {
}

func (p MasteriaCM2D) Name() string {
	return "Masteria_CM2_D"
}

func (p MasteriaCM2D) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610020001, "CM2_E")
		return false
	}
	return true
}
