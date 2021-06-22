package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TDMCEnterBoss2 struct {
}

func (p TDMCEnterBoss2) Name() string {
	return "TD_MC_enterboss2"
}

func (p TDMCEnterBoss2) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(2331) {
		script.OpenNPC(l, c)(1300013)
		return false
	}

	if script.QuestCompleted(l, c)(2333) &&
		script.QuestStarted(l, c)(2331) &&
		!script.HasItem(l, c)(4001318) {
		script.SendPinkNotice(l, c)("LOST_SEAL")
		if script.CanHold(l, c)(4001318, 1) {
			script.GainItem(l, c)(4001318, 1)
		} else {
			script.SendPinkNotice(l, c)("SEAL_INVENTORY_FULL")
		}
	}

	if script.QuestCompleted(l, c)(2333) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(106021600, 1)
		return true
	}

	if script.QuestStarted(l, c)(2332) && script.HasItem(l, c)(4032388) {
		character.ForceCompleteQuestViaNPC(l)(c.CharacterId(), 2332, 1300002)
		script.SendPinkNotice(l, c)("FOUND_PRINCESS")
		script.GainExperience(l, c)(4400)

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
		script.SendPinkNotice(l, c)("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		return false
		//	}
		//}
	} else if script.QuestStarted(l, c)(2333) || (script.QuestCompleted(l, c)(2332) && !script.QuestStarted(l, c)(2333)) {
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
		script.SendPinkNotice(l, c)("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
		return false
		//	}
		//}
	} else {
		script.SendPinkNotice(l, c)("SEEMS_TO_BE_LOCKED_LONG")
		return false
	}

}
