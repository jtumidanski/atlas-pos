package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanGarden0 struct {
}

func (p EvanGarden0) Name() string {
	return "evanGarden0"
}

func (p EvanGarden0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(100030200, "east00")
	return true
}
