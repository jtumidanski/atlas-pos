package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Pianus struct {
}

func (p Pianus) Name() string {
	return "Pianus"
}

func (p Pianus) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(230040420, "out00")
	return true
}
