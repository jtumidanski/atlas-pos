package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Lpq1 struct {
}

func (p Lpq1) Name() string {
	return "lpq1"
}

func (p Lpq1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//int nextMap = 922010300
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//String avail = eim.getProperty("2stageclear")
	//if (avail == null) {
	script.SendPinkNotice(l, c)("SEAL_BLOCKING_DOOR")
	return false
	//} else {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
}
