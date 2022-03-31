package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqEnter struct {
}

func (p GlpqEnter) Name() string {
	return "glpqEnter"
}

func (p GlpqEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992041) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610030020, "out00")
		return true
	}
	processor.SendPinkNotice(l, c)("GIANT_GATE_NO_BUDGE")
	return false
}
