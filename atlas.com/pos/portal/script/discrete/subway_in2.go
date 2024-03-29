package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type SubwayIn2 struct {
}

func (p SubwayIn2) Name() string {
	return "subway_in2"
}

func (p SubwayIn2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(103000101, 3)
	return true
}
