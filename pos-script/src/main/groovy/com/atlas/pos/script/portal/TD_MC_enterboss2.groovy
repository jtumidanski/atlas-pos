package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(2331)) {
      pi.openNpc(1300013)
      return false
   }

   if (pi.isQuestCompleted(2333) && pi.isQuestStarted(2331) && !pi.hasItem(4001318)) {
      pi.sendPinkNotice("LOST_SEAL")
      if (pi.canHold(4001318)) {
         pi.gainItem(4001318, (short) 1)
      } else {
         pi.sendPinkNotice("SEAL_INVENTORY_FULL")
      }
   }

   if (pi.isQuestCompleted(2333)) {
      pi.playPortalSound()
      pi.warp(106021600, 1)
      return true
   } else if (pi.isQuestStarted(2332) && pi.hasItem(4032388)) {
      pi.forceCompleteQuest(2332, 1300002)
      pi.sendPinkNotice("FOUND_PRINCESS")
      pi.giveCharacterExp(4400)

      EventManager em = pi.getEventManager("MK_PrimeMinister")
      Optional<MapleParty> party = pi.getPlayer().getParty()
      if (party.isPresent()) {
         MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
         if (eli.size() > 0) {
            if (em.startInstance(party.get(), pi.getMap(), 1)) {
               pi.playPortalSound()
               return true
            } else {
               pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
               return false
            }
         }
      } else {
         if (em.startInstance(pi.getPlayer())) {
            pi.playPortalSound()
            return true
         } else {
            pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
            return false
         }
      }
   } else if (pi.isQuestStarted(2333) || (pi.isQuestCompleted(2332) && !pi.isQuestStarted(2333))) {
      EventManager em = pi.getEventManager("MK_PrimeMinister")

      Optional<MapleParty> party = pi.getPlayer().getParty()
      if (party.isPresent()) {
         MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
         if (eli.size() > 0) {
            if (em.startInstance(party.get(), pi.getMap(), 1)) {
               pi.playPortalSound()
               return true
            } else {
               pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
               return false
            }
         }
      } else {
         if (em.startInstance(pi.getPlayer())) {
            pi.playPortalSound()
            return true
         } else {
            pi.sendPinkNotice("BOSS_ANOTHER_PARTY_CHALLENGING_IN_CHANNEL")
            return false
         }
      }
   } else {
      pi.sendPinkNotice("SEEMS_TO_BE_LOCKED_LONG")
      return false
   }
}