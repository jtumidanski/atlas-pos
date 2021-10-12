package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DragonEyes struct {
}

func (p DragonEyes) Name() string {
	return "dragoneyes"
}

func (p DragonEyes) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestCompleted(l, c)(22012) {
		return false
	}
	character.ForceCompleteQuest(l)(c.CharacterId(), 22012)
	script.BlockPortal(l, span, c)
	return true
}
