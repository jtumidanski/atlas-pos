package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Market06 struct {
}

func (p Market06) Name() string {
	return "market06"
}

func (p Market06) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	return generic.EnterMarket(l, span, c)
}
