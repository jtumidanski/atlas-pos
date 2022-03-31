package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor6 struct {
}

func (p RienTutor6) Name() string {
	return "rienTutor6"
}

func (p RienTutor6) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.RemoveGuide(l, c)
	processor.BlockPortal(l, span, c)
	return true
}
