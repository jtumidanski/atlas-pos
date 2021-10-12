package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3367Out struct {
}

func (p Q3367Out) Name() string {
	return "q3367out"
}

func (p Q3367Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(926130100, "in01")
	return true
}
