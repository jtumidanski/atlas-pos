package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3R4pt struct {
}

func (p Party3R4pt) Name() string {
	return "party3_r4pt"
}

func (p Party3R4pt) Enter(l logrus.FieldLogger, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim.getProperty("stage4_comb") == null) {
	//	int r = Math.floor((Math.random() * 3)).intValue() + 1
	//	int s = Math.floor((Math.random() * 3)).intValue() + 1
	//
	//	eim.setProperty("stage4_comb", "" + r + s)
	//}
	//
	//int portalName = Integer.parseInt(pi.getPortal().name().substring(4, 6))
	//int cname = (eim.getProperty("stage4_comb")).toInteger()
	//
	//boolean secondPt = true
	//if (pi.getPortal().id() < 14) {
	//	cname = Math.floor(cname / 10).intValue()
	//	secondPt = false
	//}
	//
	//if ((portalName % 10) == (cname % 10)) {    //climb
	//	int nextPortal
	//	if (secondPt) {
	//		nextPortal = 1
	//	} else {
	//		nextPortal = pi.getPortal().id() + 3
	//	}
	//
	//	pi.playPortalSound(); pi.warp(pi.getMapId(), nextPortal)
	//} else {    //fail
	//	pi.playPortalSound(); pi.warp(pi.getMapId(), 2)
	//}

	return true
}
