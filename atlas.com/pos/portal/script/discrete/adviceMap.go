package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AdviceMap struct {
}

func (p AdviceMap) Name() string {
	return "adviceMap"
}

func (p AdviceMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Press the #e#b[Up]#k arrow#n to use the portal and move to the next map.", 230, 5)
	return true
}
