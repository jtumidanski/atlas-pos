package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type Party6Stage515 struct {
}

func (p Party6Stage515) Name() string {
	return "party6_stage515"
}

func (p Party6Stage515) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if rand.Float64() < 0.1 {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(930000300, "16st")
		return true
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(930000300, "01st")
	return true
}
