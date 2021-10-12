package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MCRevive5 struct {
}

func (p MCRevive5) Name() string {
	return "MCRevive5"
}

func (p MCRevive5) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(980000501, 0)
	return true
}
