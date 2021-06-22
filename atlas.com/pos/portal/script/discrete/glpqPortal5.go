package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal5 struct {
}

func (p GlpqPortal5) Name() string {
	return "glpqPortal5"
}

func (p GlpqPortal5) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	if (eim.getIntProperty("glpq5") < 5) {
	//		pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
	//		return false
	//	} else {
	//		pi.playPortalSound()
	//		pi.warp(610030600, 0)
	//		return true
	//	}
	//}
	//
	return false
}
