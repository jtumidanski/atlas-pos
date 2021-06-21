package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type BedroomOut struct {
}

func (p BedroomOut) Name() string {
	return "bedroom_out"
}

func (p BedroomOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(2570) {
		character.PlayPortalSound(l)
		character.WarpById(l, c)(120000101, 0)
		return true
	}
	character.EarnTitle(l, c)("You still got some stuff to take care of. I can see it in your eyes. Wait...no, those are eye boogers.")
	return false
}
