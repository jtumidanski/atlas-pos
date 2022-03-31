package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM2E struct {
}

func (p MasteriaCM2E) Name() string {
	return "Masteria_CM2_E"
}

func (p MasteriaCM2E) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992039) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610020001, "CM2_F")
		return false
	}
	return true
}
