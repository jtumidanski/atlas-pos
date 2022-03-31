package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4ResurOut struct {
}

func (p S4ResurOut) Name() string {
	return "s4resur_out"
}

func (p S4ResurOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(6134) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(220070400, 3)
		return true
	}
	if !processor.CanHold(l, c)(4031448, 1) {
		processor.SendPinkNotice(l, c)("MAKE_ROOM_FOR_QUEST_ITEM")
		return false
	}
	processor.GainItem(l, c)(4031448, 1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(220070400, 3)
	return true
}
