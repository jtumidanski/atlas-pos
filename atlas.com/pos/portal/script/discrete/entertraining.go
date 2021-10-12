package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterTraining struct {
}

func (p EnterTraining) Name() string {
	return "entertraining"
}

func (p EnterTraining) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(1041) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(1010100, 4)
	} else if script.QuestStarted(l, c)(1042) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(1010200, 4)
	} else if script.QuestStarted(l, c)(1043) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(1010300, 4)
	} else if script.QuestStarted(l, c)(1044) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(1010400, 4)
	} else {
		script.SendPinkNotice(l, c)("MAI_TRAINING_REQUIREMENT")
		return false
	}
	return true
}
