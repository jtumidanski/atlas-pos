package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MoveNext struct {
}

func (p MoveNext) Name() string {
	return "moveNext"
}

func (p MoveNext) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(c.MapId() + 10, "east00")
	return true
}
