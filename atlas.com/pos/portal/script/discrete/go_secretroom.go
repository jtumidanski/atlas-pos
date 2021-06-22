package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GoSecretRoom struct {
}

func (p GoSecretRoom) Name() string {
	return "go_secretroom"
}

func (p GoSecretRoom) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestCompleted(l, c)(2335) &&
		!script.QuestStarted(l, c)(2335) &&
		script.HasItem(l, c)(4032405) {
		script.SendPinkNotice(l, c)("DOOR_LOCKED")
		return false
	}

	if script.QuestStarted(l, c)(2335) {
		character.ForceCompleteQuestViaNPC(l)(c.CharacterId(), 2335, 1300002)
		character.GainExperience(l)(c.CharacterId(), 5000)
		script.GainItem(l, c)(4032405, -1)
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(106021001, 1)
	return true
}
