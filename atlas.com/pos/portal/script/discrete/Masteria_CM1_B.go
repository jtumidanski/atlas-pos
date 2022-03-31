package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1B struct {
}

func (p MasteriaCM1B) Name() string {
	return "Masteria_CM1_B"
}

func (p MasteriaCM1B) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992039) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610020000, "CM1_C")
		return false
	}
	return true
}
