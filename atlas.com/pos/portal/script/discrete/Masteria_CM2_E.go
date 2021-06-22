package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2E struct {
}

func (p MasteriaCM2E) Name() string {
	return "Masteria_CM2_E"
}

func (p MasteriaCM2E) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610020001, "CM2_F")
		return false
	}
	return true
}
