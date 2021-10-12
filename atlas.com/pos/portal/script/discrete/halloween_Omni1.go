package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HalloweenOmni1 struct {
}

func (p HalloweenOmni1) Name() string {
	return "halloween_Omni1"
}

func (p HalloweenOmni1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.SendPinkNotice(l, c)("SEEMS_TO_BE_LOCKED")
	return true
}
