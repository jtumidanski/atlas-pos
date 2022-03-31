package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Lpq2 struct {
}

func (p Lpq2) Name() string {
	return "lpq2"
}

func (p Lpq2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//int nextMap = 922010400
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(nextMap)
	//MaplePortal targetPortal = target.getPortal("st00")
	//// only let people through if the eim is ready
	//String avail = eim.getProperty("3stageclear")
	//if (avail == null) {
	processor.SendPinkNotice(l, c)("SEAL_BLOCKING_DOOR")
	return false
	//} else {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
}
