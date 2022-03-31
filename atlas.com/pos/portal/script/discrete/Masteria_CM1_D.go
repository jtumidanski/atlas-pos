package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1D struct {
}

func (p MasteriaCM1D) Name() string {
	return "Masteria_CM1_D"
}

func (p MasteriaCM1D) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992039) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610020000, "CM1_E")
		return false
	}
	return true
}
