package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterAchter struct {
}

func (p EnterAchter) Name() string {
	return "enterAchter"
}

func (p EnterAchter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(100000201, "out02")
	return true
}
