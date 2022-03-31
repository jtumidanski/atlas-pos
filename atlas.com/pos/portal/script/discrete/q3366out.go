package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3366Out struct {
}

func (p Q3366Out) Name() string {
	return "q3366out"
}

func (p Q3366Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(926130100, "in00")
	return true
}
