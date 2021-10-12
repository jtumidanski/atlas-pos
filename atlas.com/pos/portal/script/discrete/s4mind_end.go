package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4MindEnd struct {
}

func (p S4MindEnd) Name() string {
	return "s4mind_end"
}

func (p S4MindEnd) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (!pi.getEventInstance().isEventCleared()) {
	script.SendPinkNotice(l, c)("COMPLETE_MISSION_BEFORE_PROCEEDING")
	return false
	//} else {
	//	if (pi.isQuestStarted(6410)) {
	//		pi.setQuestProgress(6410, 6411, "p2")
	//	}
	//
	//	pi.playPortalSound()
	//	pi.warp(925010400)
	//	return true
	//}
}
