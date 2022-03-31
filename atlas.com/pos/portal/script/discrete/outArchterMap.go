package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutArchterMap struct {
}

func (p OutArchterMap) Name() string {
	return "outArchterMap"
}

func (p OutArchterMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(100000000, "Achter00")
	return true
}
