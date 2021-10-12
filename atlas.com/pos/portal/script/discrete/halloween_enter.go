package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HalloweenEnter struct {
}

func (p HalloweenEnter) Name() string {
	return "halloween_enter"
}

func (p HalloweenEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(682000100, "st00")
	return true
}
