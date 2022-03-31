package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice05 struct {
}

func (p Advice05) Name() string {
	return "advice05"
}

func (p Advice05) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowInstruction(l, span, c)("Press #e#b[Q]#k#n to view the Quest window.", 250, 5)
	return true
}
