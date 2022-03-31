package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ThiefIn1 struct {
}

func (p ThiefIn1) Name() string {
	return "thief_in1"
}

func (p ThiefIn1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPCWithScript(l, c)(1063011, "ThiefPassword")
	return false
}
