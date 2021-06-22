package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Jnr2Out struct {
}

func (p Jnr2Out) Name() string {
	return "jnr2_out"
}

func (p Jnr2Out) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if (pi.getEventInstance().getIntProperty("statusStg3") == 3) {
	//	pi.playPortalSound()
	//	pi.warp(926110200, 0) //next
	//	return true
	//} else {
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
