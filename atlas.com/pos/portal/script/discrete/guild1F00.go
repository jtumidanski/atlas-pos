package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Guild1F00 struct {
}

func (p Guild1F00) Name() string {
	return "guild1F00"
}

func (p Guild1F00) Enter(l logrus.FieldLogger, c script.Context) bool {
	//int[] backPortals = [6, 8, 9, 11]
	//int idx = pi.getEventInstance().gridCheck(pi.getPlayer())
	//
	script.PlayPortalSound(l, c)
	//pi.warp(990000600, backPortals[idx])
	return true
}
