package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party6Out struct {
}

func (p Party6Out) Name() string {
	return "party6_out"
}

func (p Party6Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//
	//if (eim.isEventCleared()) {
	//	if(pi.isEventLeader()) {
	//		pi.playPortalSound()
	//		eim.warpEventTeam(930000800)
	//		return true
	//	} else {
	//		pi.sendPinkNotice("WAIT_FOR_LEADER")
	//		return false
	//	}
	//} else {
	script.SendPinkNotice(l, c)("ELIMINATE_POISON_GOLEM")
	return false
	//}
}
