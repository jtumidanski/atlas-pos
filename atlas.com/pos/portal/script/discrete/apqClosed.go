package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ApqClosed struct {
}

func (p ApqClosed) Name() string {
	return "apqClosed"
}

func (p ApqClosed) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SendPinkNotice(l, c)("GATE_IS_NOT_YET_OPENED")
	return false
}
