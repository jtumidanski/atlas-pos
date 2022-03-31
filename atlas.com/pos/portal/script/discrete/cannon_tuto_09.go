package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CannonTuto09 struct {
}

func (p CannonTuto09) Name() string {
	return "cannon_tuto_09"
}

func (p CannonTuto09) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPCWithScript(l, c)(8, "1096005")
	return true
}
