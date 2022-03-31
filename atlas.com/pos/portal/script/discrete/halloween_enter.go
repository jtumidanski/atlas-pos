package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HalloweenEnter struct {
}

func (p HalloweenEnter) Name() string {
	return "halloween_enter"
}

func (p HalloweenEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(682000100, "st00")
	return true
}
