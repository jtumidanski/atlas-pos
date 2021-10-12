package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type Party6Stage514 struct {
}

func (p Party6Stage514) Name() string {
	return "party6_stage514"
}

func (p Party6Stage514) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if rand.Float64() < 0.1 {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(930000300, "16st")
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(930000300, "15st")
	return true
}
