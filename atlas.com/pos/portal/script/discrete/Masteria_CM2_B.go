package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2B struct {
}

func (p MasteriaCM2B) Name() string {
	return "Masteria_CM2_B"
}

func (p MasteriaCM2B) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610020001, "CM2_C")
		return false
	}
	return true
}
