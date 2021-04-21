package script

import (
	"github.com/sirupsen/logrus"
)

type babyPigOut struct {
}

func BabyPigOut() PortalScript {
	return babyPigOut{}
}

func (a babyPigOut) Name() string {
	return "babyPigOut"
}

func (a babyPigOut) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.QuestCompleted(22015) {
		p.PlayPortalSound()
		p.WarpById(100030300, 2)
	} else {
		p.SendPinkNotice("RESCUE_BABY_PIG")
	}
	return true
}
