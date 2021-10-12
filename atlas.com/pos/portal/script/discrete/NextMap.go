package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type NextMap struct {
}

func (p NextMap) Name() string {
	return "NextMap"
}

func (p NextMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(c.MapId() + 100, 0)
	return true
}
