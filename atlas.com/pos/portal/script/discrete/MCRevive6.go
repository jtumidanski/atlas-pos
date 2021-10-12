package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MCRevive6 struct {
}

func (p MCRevive6) Name() string {
	return "MCRevive6"
}

func (p MCRevive6) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(980000601, 0)
	return true
}
