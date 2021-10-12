package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HighPosition struct {
}

func (p HighPosition) Name() string {
	return "highposition"
}

func (p HighPosition) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.RunMapScript(l, c)
	return false
}
