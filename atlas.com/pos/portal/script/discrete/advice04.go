package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice04 struct {
}

func (p Advice04) Name() string {
	return "advice04"
}

func (p Advice04) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowInstruction(l, span, c)("Click \r\\#b<Sera>", 100, 5)
	return true
}
