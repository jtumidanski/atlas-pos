package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCEnterBoss2 struct {
}

func (p TDMCEnterBoss2) Name() string {
	return "TD_MC_enterboss2"
}

func (p TDMCEnterBoss2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(2331) {
		processor.OpenNPC(l, c)(1300013)
		return false
	}

	if processor.QuestCompleted(l, c)(2333) &&
		processor.QuestStarted(l, c)(2331) &&
		!processor.HasItem(l, c)(4001318) {
		processor.SendPinkNotice(l, c)("LOST_SEAL")
		if processor.CanHold(l, c)(4001318, 1) {
			processor.GainItem(l, c)(4001318, 1)
		} else {
			processor.SendPinkNotice(l, c)("SEAL_INVENTORY_FULL")
		}
	}

	if processor.QuestCompleted(l, c)(2333) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106021600, 1)
		return true
	}

	if processor.QuestStarted(l, c)(2332) && processor.HasItem(l, c)(4032388) {
		character.ForceCompleteQuestViaNPC(l)(c.CharacterId(), 2332, 1300002)
		processor.SendPinkNotice(l, c)("FOUND_PRINCESS")
		processor.GainExperience(l, c)(4400)

		//EventManager em = pi.getEventManager("MK_PrimeMinister")
		//Optional<MapleParty> party = pi.getPlayer().getParty()
		//if (party.isPresent()) {
		//	MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
		//	if (eli.size() > 0) {
		//		if (em.startInstance(party.get(), pi.getMap(), 1)) {
		//			pi.playPortalSound()
		//			return true
		//		} else {
		//			pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		//			return false
		//		}
		//	}
		//} else {
		//	if (em.startInstance(pi.getPlayer())) {
		//		pi.playPortalSound()
		//		return true
		//	} else {
		processor.SendPinkNotice(l, c)("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		return false
		//	}
		//}
	} else if processor.QuestStarted(l, c)(2333) || (processor.QuestCompleted(l, c)(2332) && !processor.QuestStarted(l, c)(2333)) {
		//EventManager em = pi.getEventManager("MK_PrimeMinister")
		//
		//Optional<MapleParty> party = pi.getPlayer().getParty()
		//if (party.isPresent()) {
		//	MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
		//	if (eli.size() > 0) {
		//		if (em.startInstance(party.get(), pi.getMap(), 1)) {
		//			pi.playPortalSound()
		//			return true
		//		} else {
		//			pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		//			return false
		//		}
		//	}
		//} else {
		//	if (em.startInstance(pi.getPlayer())) {
		//		pi.playPortalSound()
		//		return true
		//	} else {
		processor.SendPinkNotice(l, c)("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		return false
		//	}
		//}
	} else {
		processor.SendPinkNotice(l, c)("SEEMS_TO_BE_LOCKED_LONG")
		return false
	}

}
