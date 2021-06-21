package discrete

import (
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
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(100030300, 2)
	} else {
		script.SendPinkNotice(l, c)("RESCUE_BABY_PIG")
	}
	return true
}
