package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice09 struct {
}

func (p Advice09) Name() string {
	return "advice09"
}

func (p Advice09) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Press #e#b[Down]#k on the arrow key#n and#e#b[Alt]#k#n at the same time to jump downwards.", 450, 6)
	return true
}
