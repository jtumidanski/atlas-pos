package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCFirst struct {
}

func (p TDMCFirst) Name() string {
	return "TD_MC_first"
}

func (p TDMCFirst) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestCompleted(l, c)(2260) ||
		script.QuestStarted(l, c)(2300) ||
		script.QuestStarted(l, c)(2301) ||
		script.QuestStarted(l, c)(2302) ||
		script.QuestStarted(l, c)(2303) ||
		script.QuestStarted(l, c)(2304) ||
		script.QuestStarted(l, c)(2305) ||
		script.QuestStarted(l, c)(2306) ||
		script.QuestStarted(l, c)(2307) ||
		script.QuestStarted(l, c)(2308) ||
		script.QuestStarted(l, c)(2309) ||
		script.QuestStarted(l, c)(2310) ||
		script.QuestCompleted(l, c)(2300) ||
		script.QuestCompleted(l, c)(2301) ||
		script.QuestCompleted(l, c)(2302) ||
		script.QuestCompleted(l, c)(2303) ||
		script.QuestCompleted(l, c)(2304) ||
		script.QuestCompleted(l, c)(2305) ||
		script.QuestCompleted(l, c)(2306) ||
		script.QuestCompleted(l, c)(2307) ||
		script.QuestCompleted(l, c)(2308) ||
		script.QuestCompleted(l, c)(2309) ||
		script.QuestCompleted(l, c)(2310) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(106020000, 0)
		return true
	}
	script.SendPinkNotice(l, c)("STRANGE_FORCE_2")
	return false
}
