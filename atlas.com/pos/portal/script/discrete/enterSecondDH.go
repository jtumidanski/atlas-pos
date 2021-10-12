package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type EnterSecondDH struct {
}

func (p EnterSecondDH) Name() string {
	return "enterSecondDH"
}

func (p EnterSecondDH) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(20201) || script.QuestStarted(l, c)(20202) || script.QuestStarted(l, c)(20203) || script.QuestStarted(l, c)(20204) || script.QuestStarted(l, c)(20205) {
		character.RemoveAll(l)(c.CharacterId(), 4032096)
		character.RemoveAll(l)(c.CharacterId(), 4032097)
		character.RemoveAll(l)(c.CharacterId(), 4032098)
		character.RemoveAll(l)(c.CharacterId(), 4032099)
		character.RemoveAll(l)(c.CharacterId(), 4032100)

		mapIds := []uint32{108000600, 108000601, 108000602}
		random := rand.Intn(len(mapIds))
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(mapIds[random], 0)
		return true
	}
	return false
}
