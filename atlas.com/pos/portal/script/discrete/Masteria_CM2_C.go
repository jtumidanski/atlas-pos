package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2C struct {
}

func (p MasteriaCM2C) Name() string {
	return "Masteria_CM2_C"
}

func (p MasteriaCM2C) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610020001, "CM2_D")
		return false
	}
	return true
}
