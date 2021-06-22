package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RienTutor8 struct {
}

func (p RienTutor8) Name() string {
	return "rienTutor8"
}

func (p RienTutor8) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.IsJob(l)(c.CharacterId(), job.Legend) {
		if script.QuestStarted(l, c)(21015) {
			script.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
		if script.QuestStarted(l, c)(21016) {
			script.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
		if script.QuestStarted(l, c)(21017) {
			script.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(140010000, 2)
	return true
}
