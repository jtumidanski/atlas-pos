package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Market02 struct {
}

func (p Market02) Name() string {
	return "market02"
}

func (p Market02) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	return generic.EnterMarket(l, span, c)
}
