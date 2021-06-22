package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Jnr4R4 struct {
}

func (p Jnr4R4) Name() string {
	return "jnr4_r4"
}

func (p Jnr4R4) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//int area = eim.getIntProperty("statusStg5")
	//int reg = 3
	//
	//if((area >> reg) % 2 == 0) {
	//	area |= (1 << reg)
	//	eim.setIntProperty("statusStg5", area)
	//
	//	pi.playPortalSound(); pi.warp(926110301 + reg, 0) //next
	//	return true
	//} else {
	script.SendPinkNotice(l, c)("ROOM_ALREADY_BEING_EXPLORED")
	return false
	//}
}
