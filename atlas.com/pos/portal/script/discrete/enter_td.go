package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterTd struct {
}

func (p EnterTd) Name() string {
	return "enter_td"
}

func (p EnterTd) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(600000000, "yn00")
	return true
}
