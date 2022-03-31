package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutPerrion1 struct {
}

func (p OutPerrion1) Name() string {
	return "outPerrion_1"
}

func (p OutPerrion1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SendPinkNotice(l, c)("TEMPLE_SHORTCUT")
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(105100000, 2)
	return true
}
