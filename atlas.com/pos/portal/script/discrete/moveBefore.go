package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MoveBefore struct {
}

func (p MoveBefore) Name() string {
	return "moveBefore"
}

func (p MoveBefore) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(c.MapId() - 10, "west00")
	return true
}
