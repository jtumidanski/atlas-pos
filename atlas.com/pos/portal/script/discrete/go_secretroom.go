package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GoSecretRoom struct {
}

func (p GoSecretRoom) Name() string {
	return "go_secretroom"
}

func (p GoSecretRoom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestCompleted(l, c)(2335) &&
		!processor.QuestStarted(l, c)(2335) &&
		processor.HasItem(l, c)(4032405) {
		processor.SendPinkNotice(l, c)("DOOR_LOCKED")
		return false
	}

	if processor.QuestStarted(l, c)(2335) {
		character.ForceCompleteQuestViaNPC(l)(c.CharacterId(), 2335, 1300002)
		character.GainExperience(l)(c.CharacterId(), 5000)
		processor.GainItem(l, c)(4032405, -1)
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(106021001, 1)
	return true
}
