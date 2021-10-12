package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterNautil struct {
}

func (p EnterNautil) Name() string {
	return "enter_nautil"
}

func (p EnterNautil) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(120010000, "nt01")
	return true
}
