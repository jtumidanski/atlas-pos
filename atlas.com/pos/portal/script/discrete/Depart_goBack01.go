package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartGoBack01 struct {
}

func (p DepartGoBack01) Name() string {
	return "DepartBack01"
}

func (p DepartGoBack01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(c.MapId() - 10, "left01")
	return true
}
