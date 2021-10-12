package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor6 struct {
}

func (p RienTutor6) Name() string {
	return "rienTutor6"
}

func (p RienTutor6) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.RemoveGuide(l, c)
	script.BlockPortal(l, span, c)
	return true
}
