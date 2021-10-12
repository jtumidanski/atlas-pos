package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MCRevive2 struct {
}

func (p MCRevive2) Name() string {
	return "MCRevive2"
}

func (p MCRevive2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(980000201, 0)
	return true
}
