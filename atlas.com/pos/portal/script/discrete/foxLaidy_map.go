package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type FoxLaidyMap struct {
}

func (p FoxLaidyMap) Name() string {
	return "foxLaidy_map"
}

func (p FoxLaidyMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(3647) && script.HasItem(l, c)(4031793) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(222010200, "east00")
		return true
	}

	if !script.QuestStarted(l, c)(23647) {
		character.ForceStartQuest(l)(c.CharacterId(), 23647)
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(922220000, "east00")
	return true
}
