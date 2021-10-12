package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice01 struct {
}

func (p Advice01) Name() string {
	return "advice01"
}

func (p Advice01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Click \r\\#b<Heena>#k", 100, 5)
	return true
}
