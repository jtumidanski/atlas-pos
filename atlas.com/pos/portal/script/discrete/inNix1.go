package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InNix1 struct {
}

func (p InNix1) Name() string {
	return "inNix1"
}

func (p InNix1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(240020600,"out00")
	return true
}
