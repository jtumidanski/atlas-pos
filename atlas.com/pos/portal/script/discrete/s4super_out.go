package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type S4SuperOut struct {
}

func (p S4SuperOut) Name() string {
	return "s4super_out"
}

func (p S4SuperOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	//int exit = pi.getEventInstance().getIntProperty("canLeave")
	//if (exit == 0) {
	//	pi.sendPinkNotice("WAIT_TO_LEAVE")
	//	return false
	//} else if (exit == 2) {
	//	pi.playPortalSound()
	//	pi.warp(912010200)
	//	return true
	//} else {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(120000101)
	return true
	//}
}
