package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Market13 struct {
}

func (p Market13) Name() string {
	return "market13"
}

func (p Market13) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	return generic.EnterMarket(l, span, c)
}
