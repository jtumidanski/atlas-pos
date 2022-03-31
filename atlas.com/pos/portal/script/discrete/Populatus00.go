package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Populatus00 struct {
}

func (p Populatus00) Name() string {
	return "Populatus00"
}

func (p Populatus00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !((processor.QuestStarted(l, c)(6361) && processor.HasItem(l, c)(4031870)) || processor.QuestCompleted(l, c)(6361)) && !processor.QuestCompleted(l, c)(6363) {
		//EventManager em = pi.getEventManager("PapulatusBattle")
		//
		//if (pi.getParty() == null) {
		//	pi.sendPinkNotice("BOSS_PARTY_NEEDED")
		//	return false
		//} else if (!pi.isLeader()) {
		//	pi.sendPinkNotice("BOSS_PARTY_LEADER_START")
		//	return false
		//} else {
		//	MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
		//	if (eli.size() > 0) {
		//		if (!em.startInstance(pi.getParty().orElseThrow(), pi.getPlayer().getMap(), 1)) {
		//			pi.sendPinkNotice("BOSS_ALREADY_STARTED")
		//			return false
		//		}
		//	} else {  //this should never appear
		//		pi.sendPinkNotice("BOSS_CANNOT_START_YET")
		//		return false
		//	}

		processor.PlayPortalSound(l, c)
		return true
		//}
	} else {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(922020300, 0)
		return true
	}
}
