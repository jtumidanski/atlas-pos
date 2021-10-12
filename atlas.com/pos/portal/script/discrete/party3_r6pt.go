package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3R6pt struct {
}

func (p Party3R6pt) Name() string {
	return "party3_r6pt"
}

func (p Party3R6pt) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim.getProperty("stage6_comb") == null) {
	//	String comb = "0"
	//
	//	for (int i = 0; i < 16; i++) {
	//		int r = Math.floor((Math.random() * 4)).intValue() + 1
	//		comb += r.toString()
	//	}
	//
	//	eim.setProperty("stage6_comb", comb)
	//}
	//
	//String comb = eim.getProperty("stage6_comb")
	//
	//String name = pi.getPortal().name().substring(2, 5)
	//int portalId = Integer.parseInt(name)
	//
	//
	//int pRow = (int) Math.floor(portalId / 10)
	//int pCol = portalId % 10
	//
	//if (pCol == Integer.parseInt(comb.substring(pRow, pRow + 1))) {    //climb
	//	pi.playPortalSound()
	//	pi.warp(pi.getMapId(), (pRow % 4 != 0) ? pi.getPortal().id() + 4 : ((int) (pRow / 4)))
	//} else {    //fail
	//	pRow--
	//	pi.playPortalSound()
	//	pi.warp(pi.getMapId(), (pRow / 4) > 1 ? ((int) (pRow / 4)) : 5)
	//}

	return true
}
