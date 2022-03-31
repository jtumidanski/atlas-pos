package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterBackStreet struct {
}

func (p EnterBackStreet) Name() string {
	return "enterBackStreet"
}

func (p EnterBackStreet) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestActive(l, c)(21747) || processor.QuestActive(l, c)(21744) && processor.QuestCompleted(l, c)(21745) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(925040000, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("NO_PERMISSION_TO_ENTER")
	return false
}
