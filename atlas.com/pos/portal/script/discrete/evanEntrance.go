package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanEntrance struct {
}

func (p EvanEntrance) Name() string {
	return "evanEntrance"
}

func (p EvanEntrance) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(100030400, "east00")
	return true
}
