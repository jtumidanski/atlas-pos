package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CurseForest struct {
}

func (p CurseForest) Name() string {
	return "curseforest"
}

func (p CurseForest) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(2224) || script.QuestStarted(l, c)(2226) || script.QuestCompleted(l, c)(2227) {
		hd := script.GetHourOfDay(l)
		if !((hd >= 0 && hd < 7) || hd >= 17) {
			character.SendPinkNotice(l, c)("CURSED_FOREST_NOT_NOW")
			return false
		} else {
			character.PlayPortalSound(l)
			var mid = uint32(0)
			if script.QuestCompleted(l, c)(2227) {
				mid = 910100001
			} else {
				mid = 910100000
			}
			character.WarpByName(l, c)(mid, "out00")
			return true
		}
	}

	character.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
