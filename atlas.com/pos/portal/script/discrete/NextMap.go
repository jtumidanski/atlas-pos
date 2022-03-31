package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type NextMap struct {
}

func (p NextMap) Name() string {
	return "NextMap"
}

func (p NextMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(c.MapId() + 100, 0)
	return true
}
