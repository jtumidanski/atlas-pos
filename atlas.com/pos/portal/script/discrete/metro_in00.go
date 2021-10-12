package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MetroIn00 struct {
}

func (p MetroIn00) Name() string {
	return "metro_in00"
}

func (p MetroIn00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.OpenNPC(l, c)(1052115)
	return true
}
