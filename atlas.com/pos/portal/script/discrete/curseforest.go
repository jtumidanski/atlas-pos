package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CurseForest struct {
}

func (p CurseForest) Name() string {
	return "curseforest"
}

func (p CurseForest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(2224) || script.QuestStarted(l, c)(2226) || script.QuestCompleted(l, c)(2227) {
		hd := script.GetHourOfDay(l)
		if !((hd >= 0 && hd < 7) || hd >= 17) {
			script.SendPinkNotice(l, c)("CURSED_FOREST_NOT_NOW")
			return false
		} else {
			script.PlayPortalSound(l, c)
			var mid = uint32(0)
			if script.QuestCompleted(l, c)(2227) {
				mid = 910100001
			} else {
				mid = 910100000
			}
			script.WarpByName(l, span, c)(mid, "out00")
			return true
		}
	}

	script.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
