package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Gryphius struct {
}

func (p Gryphius) Name() string {
	return "gryphius"
}

func (p Gryphius) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(240020101, "out00")
	return true
}
