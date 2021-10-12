package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3368Out struct {
}

func (p Q3368Out) Name() string {
	return "q3368out"
}

func (p Q3368Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(926130100, "in02")
	return true
}
