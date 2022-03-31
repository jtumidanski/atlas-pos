package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2B struct {
}

func (p MasteriaCM2B) Name() string {
	return "Masteria_CM2_B"
}

func (p MasteriaCM2B) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992039) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610020001, "CM2_C")
		return false
	}
	return true
}
