package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Obstacle struct {
}

func (p Obstacle) Name() string {
	return "obstacle"
}

func (p Obstacle) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(100202) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106020400, 2)
		return true
	}
	if processor.HasItem(l, c)(4000507) {
		processor.GainItem(l, c)(4000507, -1)
		processor.SendPinkNotice(l, c)("POISON_SPORE_USED")
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106020400, 2)
		return true
	}
	processor.SendPinkNotice(l, c)("OVERGROWN_VINES")
	return false
}
