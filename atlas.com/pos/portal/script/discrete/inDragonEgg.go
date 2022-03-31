package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InDragonEgg struct {
}

func (p InDragonEgg) Name() string {
	return "inDrnEgg"
}

func (p InDragonEgg) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	if processor.QuestStarted(l, c)(22005) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(900020100, 0)
		return true
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(100030301, 0)
	return true
}
