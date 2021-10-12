package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj5Rp struct {
}

func (p Rnj5Rp) Name() string {
	return "rnj5_rp"
}

func (p Rnj5Rp) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//String mapPlayer = "stage6_comb" + (pi.getMapId() % 10)
	//EventInstanceManager eim = pi.getEventInstance()
	//
	//String comb = eim.getProperty(mapPlayer)
	//
	//String name = pi.getPortal().name().substring(2, 4)
	//int portalId = Integer.parseInt(name)
	//
	//
	//int pRow = (int) Math.floor(portalId / 10)
	//int pCol = (portalId % 10)
	//
	//if (pCol == Integer.parseInt(comb.substring(pRow, pRow + 1))) {    //climb
	//	if (pRow < 9) {
	//		pi.playPortalSound(); pi.warp(pi.getMapId(), pi.getPortal().id() + 4)
	//	} else {
	//		if (eim.getIntProperty("statusStg6") == 0) {
	//			eim.setIntProperty("statusStg6", 1)
	//			eim.giveEventPlayersStageReward(6)
	//		}
	//
	//		pi.playPortalSound()
	//		pi.warp(pi.getMapId(), 1)
	//	}
	//
	//} else {    //fail
	//	pi.playPortalSound()
	//	pi.warp(pi.getMapId(), 2)
	//}

	return true
}
