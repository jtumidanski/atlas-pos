package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Lpq7 struct {
}

func (p Lpq7) Name() string {
	return "lpq7"
}

func (p Lpq7) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//int nextMap = 922010900
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//// only let people through if the eim is ready
	//String avail = eim.getProperty("8stageclear")
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
