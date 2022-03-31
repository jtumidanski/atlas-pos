package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice07 struct {
}

func (p Advice07) Name() string {
	return "advice07"
}

func (p Advice07) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowInstruction(l, span, c)("You can view the World Map by pressing the #e#b[W]#k#nkey.", 350, 5)
	return true
}
