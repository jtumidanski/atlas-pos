package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Lpq4 struct {
}

func (p Lpq4) Name() string {
	return "lpq4"
}

func (p Lpq4) Enter(l logrus.FieldLogger, c script.Context) bool {
	//int nextMap = 922010600
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//// only let people through if the eim is ready
	//String avail = eim.getProperty("5stageclear")
	//if (avail == null) {
	//	// can't go thru eh?
	script.SendPinkNotice(l, c)("SEAL_BLOCKING_DOOR")
	return false
	//} else {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
}
