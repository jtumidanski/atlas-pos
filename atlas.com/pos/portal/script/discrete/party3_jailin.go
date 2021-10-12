package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3JailIn struct {
}

func (p Party3JailIn) Name() string {
	return "party3_jailin"
}

func (p Party3JailIn) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//MapleMap map = pi.getMap()
	//
	//int jailn = (pi.getMap().getId() / 10) % 10
	//int maxToggles = (jailn == 1) ? 7 : 6
	//
	//String mapProp = pi.getEventInstance().getProperty("jail" + jailn)
	//
	//if (mapProp == null) {
	//	int seq = 0
	//
	//	for (int i = 1; i <= maxToggles; i++) {
	//		if (Math.random() < 0.5) {
	//			seq += (1 << i)
	//		}
	//	}
	//
	//	pi.getEventInstance().setProperty("jail" + jailn, seq)
	//	mapProp = seq
	//}
	//
	//int mapProp2 = (mapProp).toInteger()
	//if (mapProp2 != 0) {
	//	int countMiss = 0
	//	for (int i = 1; i <= maxToggles; i++) {
	//		if (!(pi.getMap().getReactorByName("lever" + i).getState() == ((mapProp2 >> i) % 2).byteValue())) {
	//			countMiss++
	//		}
	//	}
	//
	//	if (countMiss > 0) {
	//		MasterBroadcaster.getInstance().sendToAllInMap(map, new ShowEffect("quest/party/wrong_kor"))
	//		MasterBroadcaster.getInstance().sendToAllInMap(map, new PlaySound("Party1/Failed"))
	//
	//		pi.sendPinkNotice("LEVERS_MISPLACED", countMiss)
	//		return false
	//	}
	//
	//	MasterBroadcaster.getInstance().sendToAllInMap(map, new ShowEffect("quest/party/clear"))
	//	MasterBroadcaster.getInstance().sendToAllInMap(map, new PlaySound("Party1/Clear"))
	//	pi.getEventInstance().setProperty("jail" + jailn, "0")
	//}
	//
	//pi.playPortalSound(); pi.warp(pi.getMapId() + 2, 0)
	return true
}
