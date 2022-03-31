package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Kpq0 struct {
}

func (p Kpq0) Name() string {
	return "kpq0"
}

func (p Kpq0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getPlayer().getEventInstance()
	//MapleMap target = eim.getMapInstance(103000801)
	//
	//if (eim.getProperty("1stageclear") != null) {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, target.getPortal("st00"))
	//	return true
	//}
	//else {
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
	//}
}
