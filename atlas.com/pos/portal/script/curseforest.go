package script

import (
	"github.com/sirupsen/logrus"
)

type curseForest struct {
}

func CurseForest() PortalScript {
	return curseForest{}
}

func (a curseForest) Name() string {
	return "curseforest"
}

func (a curseForest) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.QuestStarted(2224) || p.QuestStarted(2226) || p.QuestCompleted(2227) {
		hd := p.GetHourOfDay()
		if !((hd >= 0 && hd < 7) || hd >= 17) {
			p.SendPinkNotice("CURSED_FOREST_NOT_NOW")
			return false
		} else {
			p.PlayPortalSound()
			var mid = uint32(0)
			if p.QuestCompleted(2227) {
				mid = 910100001
			} else {
				mid = 910100000
			}
			p.WarpByName(mid, "out00")
			return true
		}
	}

	p.SendPinkNotice("CANNOT_ACCESS")
	return false
}
