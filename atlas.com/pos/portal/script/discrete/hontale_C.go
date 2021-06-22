package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type HorntaleC struct {
}

func (p HorntaleC) Name() string {
	return "hontale_C"
}

func (p HorntaleC) Enter(l logrus.FieldLogger, c script.Context) bool {
	//if (pi.isEventLeader()) {
	//	EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//	int target
	//	byte theWay = pi.getMap().getReactorByName("light").getState()
	//	if (theWay == (byte) 1) {
	//		target = 240050300 //light
	//	} else if (theWay == (byte) 3) {
	//		target = 240050310 //dark
	//	} else {
	//		pi.sendPinkNotice("HORNTAIL_HIT_LIGHT_BULB")
	//		return false
	//	}
	//
	//	pi.playPortalSound()
	//	eim.warpEventTeam(target)
	//	return true
	//} else {
	script.SendLightBlueNotice(l, c)("HORNTAIL_PARTY_LEADER")
	return false
	//}
}
