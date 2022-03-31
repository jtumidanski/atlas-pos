package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj4R2 struct {
}

func (p Rnj4R2) Name() string {
	return "rnj4_r2"
}

func (p Rnj4R2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//int area = eim.getIntProperty("statusStg5")
	//int reg = 1
	//
	//if ((area >> reg) % 2 == 0) {
	//	area |= (1 << reg)
	//	eim.setIntProperty("statusStg5", area)
	//
	//	pi.playPortalSound()
	//	pi.warp(926100301 + reg, 0) //next
	//	return true
	//} else {
	processor.SendPinkNotice(l, c)("ROOM_ALREADY_BEING_EXPLORED")
	return false
	//}
}
