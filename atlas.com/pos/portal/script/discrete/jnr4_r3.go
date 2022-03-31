package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr4R3 struct {
}

func (p Jnr4R3) Name() string {
	return "jnr4_r3"
}

func (p Jnr4R3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//int area = eim.getIntProperty("statusStg5")
	//int reg = 2
	//
	//if((area >> reg) % 2 == 0) {
	//	area |= (1 << reg)
	//	eim.setIntProperty("statusStg5", area)
	//
	//	pi.playPortalSound(); pi.warp(926110301 + reg, 0) //next
	//	return true
	//} else {
	processor.SendPinkNotice(l, c)("ROOM_ALREADY_BEING_EXPLORED")
	return false
	//}
}
