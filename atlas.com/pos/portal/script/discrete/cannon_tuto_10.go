package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CannonTuto10 struct {
}

func (p CannonTuto10) Name() string {
	return "cannon_tuto_10"
}

func (p CannonTuto10) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.UnlockUI(l, c)
	return true
}
