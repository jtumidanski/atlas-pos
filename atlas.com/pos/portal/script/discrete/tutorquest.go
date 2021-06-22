package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TutorQuest struct {
}

func (p TutorQuest) Name() string {
	return "tutorquest"
}

func (p TutorQuest) Enter(l logrus.FieldLogger, c script.Context) bool {
	if c.MapId() == 130030001 {
		if script.QuestStarted(l, c)(20010) {
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(130030002, 0)
			return true
		}
		script.SendPinkNotice(l, c)("CLICK_ON_THE_NPC_FIRST_TO_RECEIVE")
		return false
	}

	if c.MapId() == 130030002 {
		if script.QuestCompleted(l, c)(20011) {
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(130030003, 0)
			return true
		}
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	if c.MapId() == 130030003 {
		if script.QuestCompleted(l, c)(20012) {
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(130030004, 0)
			return true
		}
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	if c.MapId() == 130030004 {
		if script.QuestCompleted(l, c)(20013) {
			script.PlayPortalSound(l, c)
			script.WarpById(l, c)(130030005, 0)
			return true
		}
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING_2")
		return false
	}

	return false
}
