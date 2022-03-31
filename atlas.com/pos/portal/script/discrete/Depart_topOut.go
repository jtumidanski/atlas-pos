package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartTopOut struct {
}

func (p DepartTopOut) Name() string {
	return "Depart_topOut"
}

func (p DepartTopOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(103040300, 1)
	return true
}
