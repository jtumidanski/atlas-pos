package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1B struct {
}

func (p MasteriaCM1B) Name() string {
	return "Masteria_CM1_B"
}

func (p MasteriaCM1B) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610020000, "CM1_C")
		return false
	}
	return true
}
