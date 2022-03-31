package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutPerrion2 struct {
}

func (p OutPerrion2) Name() string {
	return "outPerrion_2"
}

func (p OutPerrion2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(105100000, 0)
	return true
}
