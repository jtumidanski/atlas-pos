package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Market10 struct {
}

func (p Market10) Name() string {
	return "market10"
}

func (p Market10) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	return generic.EnterMarket(l, span, c)
}
