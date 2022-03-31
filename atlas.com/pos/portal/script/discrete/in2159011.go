package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type In2159011 struct {
}

func (p In2159011) Name() string {
	return "in2159011"
}

func (p In2159011) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPC(l, c)(2159011)
	return true
}
