package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GagaSuccess struct {
}

func (p GagaSuccess) Name() string {
	return "gaga_success"
}

func (p GagaSuccess) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(922240100 + (c.MapId() - 922240000), 0)
	return true
}
