package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanEntrance struct {
}

func (p EvanEntrance) Name() string {
	return "evanEntrance"
}

func (p EvanEntrance) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(100030400, "east00")
	return true
}
