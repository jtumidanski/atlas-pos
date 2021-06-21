package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanFarmCT struct {
}

func (p EvanFarmCT) Name() string {
	return "evanFarmCT"
}

func (p EvanFarmCT) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(22010) || !character.IsJob(l)(c.CharacterId(), job.Evan) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(100030310, 0)
	} else {
		script.SendPinkNotice(l, c)("CANNOT_ENTER_LUSH_FOREST_WITHOUT")
	}
	return true
}
