package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCM1C struct {
}

func (p MasteriaCM1C) Name() string {
	return "Masteria_CM1_C"
}

func (p MasteriaCM1C) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992039) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610020000, "CM1_D")
		return false
	}
	return true
}
