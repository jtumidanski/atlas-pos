package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CannonTuto06 struct {
}

func (p CannonTuto06) Name() string {
	return "cannon_tuto_06"
}

func (p CannonTuto06) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.SetDirectionStatus(l, c)(true)
	script.LockUI(l, c)
	script.OpenNPCWithScript(l, c)(3, "1096003")
	return true
}
