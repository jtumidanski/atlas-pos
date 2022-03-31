package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice02 struct {
}

func (p Advice02) Name() string {
	return "advice02"
}

func (p Advice02) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowInstruction(l, span, c)("Press #e#b[Alt]#k#n to\r\\ JUMP.", 100, 5)
	return true
}
