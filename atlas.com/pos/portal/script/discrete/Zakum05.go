package discrete

import (
	"atlas-pos/portal/script"
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
	if !(script.QuestStarted(l, c)(100200) || script.QuestCompleted(l, c)(100200)) {
		script.SendPinkNotice(l, c)("ZAKUM_MASTER_APPROVAL")
		return false
	}

	if !script.QuestCompleted(l, c)(100201) {
		script.SendPinkNotice(l, c)("ZAKUM_COMPLETE_TRIALS_2")
		return false
	}

	if !script.HasItem(l, c)(4001017) {
		script.SendPinkNotice(l, c)("ZAKUM_NEED_EYE_OF_FIRE")
		return false
	}

	if reactor.ById(l)(c.WorldId(), c.ChannelId(), c.MapId(), 2118002).State() > 0 {
		script.SendPinkNotice(l, c)("ENTRANCE_BLOCKED")
		return false
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(211042400, "west00")
	return true
}
