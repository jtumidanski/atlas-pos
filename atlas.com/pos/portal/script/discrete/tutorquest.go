package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TutorQuest struct {
}

func (p TutorQuest) Name() string {
	return "tutorquest"
}

func (p TutorQuest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() == 130030001 {
		if processor.QuestStarted(l, c)(20010) {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(130030002, 0)
			return true
		}
		processor.SendPinkNotice(l, c)("CLICK_ON_THE_NPC_FIRST_TO_RECEIVE")
		return false
	}

	if c.MapId() == 130030002 {
		if processor.QuestCompleted(l, c)(20011) {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(130030003, 0)
			return true
		}
		processor.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	if c.MapId() == 130030003 {
		if processor.QuestCompleted(l, c)(20012) {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(130030004, 0)
			return true
		}
		processor.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	if c.MapId() == 130030004 {
		if processor.QuestCompleted(l, c)(20013) {
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(130030005, 0)
			return true
		}
		processor.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	return false
}
