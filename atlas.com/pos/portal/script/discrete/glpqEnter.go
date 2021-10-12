package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqEnter struct {
}

func (p GlpqEnter) Name() string {
	return "glpqEnter"
}

func (p GlpqEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992041) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610030020, "out00")
		return true
	}
	script.SendPinkNotice(l, c)("GIANT_GATE_NO_BUDGE")
	return false
}
