package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ReundoDraco struct {
}

func (p ReundoDraco) Name() string {
	return "reundodraco"
}

func (p ReundoDraco) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.BlockPortal(l, span, c)
	return true
}
