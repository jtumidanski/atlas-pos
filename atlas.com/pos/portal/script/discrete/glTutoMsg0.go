package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlTutoMsg0 struct {
}

func (p GlTutoMsg0) Name() string {
	return "glTutoMsg0"
}

func (p GlTutoMsg0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInstruction(l, span, c)("Once you leave this area you won't be able to return.", 150, 5)
	return true
}
