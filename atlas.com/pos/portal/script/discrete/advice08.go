package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Advice08 struct {
}

func (p Advice08) Name() string {
	return "advice08"
}

func (p Advice08) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowInstruction(l, span, c)("You can check your character's stats by pressing the #e#b[S]#k#nkey.", 350, 5)
	return true
}
