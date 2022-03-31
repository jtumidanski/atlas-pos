package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MetroChat00 struct {
}

func (p MetroChat00) Name() string {
	return "metro_Chat00"
}

func (p MetroChat00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.ShowEffect(l, c)("Effect/Direction2.img/metro/Im")
	return true
}
