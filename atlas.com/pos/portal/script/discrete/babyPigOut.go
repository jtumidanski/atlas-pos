package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type BabyPigOut struct {
}

func (p BabyPigOut) Name() string {
	return "babyPigOut"
}

func (p BabyPigOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(22015) {
		character.PlayPortalSound(l)
		character.WarpById(l, c)(100030300, 2)
	} else {
		character.SendPinkNotice(l, c)("RESCUE_BABY_PIG")
	}
	return true
}
