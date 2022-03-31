package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartGoBack00 struct {
}

func (p DepartGoBack00) Name() string {
	return "DepartBack00"
}

func (p DepartGoBack00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(c.MapId() - 10, "left00")
	return true
}
