package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice03 struct {
}

func (p Advice03) Name() string {
	return "advice03"
}

func (p Advice03) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Press #e#b[Up]#k on the arrow key#n to climb up the ladder or rope.", 350, 5)
	return true
}
