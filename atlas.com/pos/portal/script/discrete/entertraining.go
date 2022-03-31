package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterTraining struct {
}

func (p EnterTraining) Name() string {
	return "entertraining"
}

func (p EnterTraining) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(1041) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(1010100, 4)
	} else if processor.QuestStarted(l, c)(1042) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(1010200, 4)
	} else if processor.QuestStarted(l, c)(1043) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(1010300, 4)
	} else if processor.QuestStarted(l, c)(1044) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(1010400, 4)
	} else {
		processor.SendPinkNotice(l, c)("MAI_TRAINING_REQUIREMENT")
		return false
	}
	return true
}
