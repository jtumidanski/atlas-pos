package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice06 struct {
}

func (p Advice06) Name() string {
	return "advice06"
}

func (p Advice06) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Press the #e#b[Up]#k arrow#n to use the portal \r\\and move to the next map.", 230, 5)
	return true
}
