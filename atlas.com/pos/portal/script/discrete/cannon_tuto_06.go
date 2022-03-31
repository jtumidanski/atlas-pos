package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CannonTuto06 struct {
}

func (p CannonTuto06) Name() string {
	return "cannon_tuto_06"
}

func (p CannonTuto06) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SetDirectionStatus(l, c)(true)
	processor.LockUI(l, c)
	processor.OpenNPCWithScript(l, c)(3, "1096003")
	return true
}
