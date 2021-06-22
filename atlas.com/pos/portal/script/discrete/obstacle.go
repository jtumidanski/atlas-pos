package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Obstacle struct {
}

func (p Obstacle) Name() string {
	return "obstacle"
}

func (p Obstacle) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(100202) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(106020400, 2)
		return true
	}
	if script.HasItem(l, c)(4000507) {
		script.GainItem(l, c)(4000507, -1)
		script.SendPinkNotice(l, c)("POISON_SPORE_USED")
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(106020400, 2)
		return true
	}
	script.SendPinkNotice(l, c)("OVERGROWN_VINES")
	return false
}
