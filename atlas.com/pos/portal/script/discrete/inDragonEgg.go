package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InDragonEgg struct {
}

func (p InDragonEgg) Name() string {
	return "inDrnEgg"
}

func (p InDragonEgg) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	if script.QuestStarted(l, c)(22005) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(900020100, 0)
		return true
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(100030301, 0)
	return true
}
