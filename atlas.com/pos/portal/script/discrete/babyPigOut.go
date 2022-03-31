package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type BabyPigOut struct {
}

func (p BabyPigOut) Name() string {
	return "babyPigOut"
}

func (p BabyPigOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(22015) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(100030300, 2)
	} else {
		processor.SendPinkNotice(l, c)("RESCUE_BABY_PIG")
	}
	return true
}
