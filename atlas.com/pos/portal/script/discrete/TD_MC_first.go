package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCFirst struct {
}

func (p TDMCFirst) Name() string {
	return "TD_MC_first"
}

func (p TDMCFirst) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(2260) ||
		processor.QuestStarted(l, c)(2300) ||
		processor.QuestStarted(l, c)(2301) ||
		processor.QuestStarted(l, c)(2302) ||
		processor.QuestStarted(l, c)(2303) ||
		processor.QuestStarted(l, c)(2304) ||
		processor.QuestStarted(l, c)(2305) ||
		processor.QuestStarted(l, c)(2306) ||
		processor.QuestStarted(l, c)(2307) ||
		processor.QuestStarted(l, c)(2308) ||
		processor.QuestStarted(l, c)(2309) ||
		processor.QuestStarted(l, c)(2310) ||
		processor.QuestCompleted(l, c)(2300) ||
		processor.QuestCompleted(l, c)(2301) ||
		processor.QuestCompleted(l, c)(2302) ||
		processor.QuestCompleted(l, c)(2303) ||
		processor.QuestCompleted(l, c)(2304) ||
		processor.QuestCompleted(l, c)(2305) ||
		processor.QuestCompleted(l, c)(2306) ||
		processor.QuestCompleted(l, c)(2307) ||
		processor.QuestCompleted(l, c)(2308) ||
		processor.QuestCompleted(l, c)(2309) ||
		processor.QuestCompleted(l, c)(2310) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106020000, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("STRANGE_FORCE_2")
	return false
}
