package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CurseForest struct {
}

func (p CurseForest) Name() string {
	return "curseforest"
}

func (p CurseForest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(2224) || processor.QuestStarted(l, c)(2226) || processor.QuestCompleted(l, c)(2227) {
		hd := processor.GetHourOfDay(l)
		if !((hd >= 0 && hd < 7) || hd >= 17) {
			processor.SendPinkNotice(l, c)("CURSED_FOREST_NOT_NOW")
			return false
		} else {
			processor.PlayPortalSound(l, c)
			var mid = uint32(0)
			if processor.QuestCompleted(l, c)(2227) {
				mid = 910100001
			} else {
				mid = 910100000
			}
			processor.WarpByName(l, span, c)(mid, "out00")
			return true
		}
	}

	processor.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
