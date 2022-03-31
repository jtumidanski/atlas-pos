package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"atlas-pos/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Zakum05 struct {
}

func (p Zakum05) Name() string {
	return "Zakum05"
}

func (p Zakum05) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !(processor.QuestStarted(l, c)(100200) || processor.QuestCompleted(l, c)(100200)) {
		processor.SendPinkNotice(l, c)("ZAKUM_MASTER_APPROVAL")
		return false
	}

	if !processor.QuestCompleted(l, c)(100201) {
		processor.SendPinkNotice(l, c)("ZAKUM_COMPLETE_TRIALS_2")
		return false
	}

	if !processor.HasItem(l, c)(4001017) {
		processor.SendPinkNotice(l, c)("ZAKUM_NEED_EYE_OF_FIRE")
		return false
	}

	if r, err := reactor.GetById(l, span)(2118002); err == nil && r.State() > 0 {
		processor.SendPinkNotice(l, c)("ENTRANCE_BLOCKED")
		return false
	}

	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(211042400, "west00")
	return true
}
