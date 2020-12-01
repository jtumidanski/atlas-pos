package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!((pi.isQuestStarted(6361) && pi.hasItem(4031870, 1)) || (pi.isQuestCompleted(6361) && !pi.isQuestCompleted(6363)))) {
      EventManager em = pi.getEventManager("PapulatusBattle")

      if (pi.getParty() == null) {
         pi.sendPinkNotice("BOSS_PARTY_NEEDED")
         return false
      } else if (!pi.isLeader()) {
         pi.sendPinkNotice("BOSS_PARTY_LEADER_START")
         return false
      } else {
         MaplePartyCharacter[] eli = em.getEligibleParty(pi.getParty().orElseThrow())
         if (eli.size() > 0) {
            if (!em.startInstance(pi.getParty().orElseThrow(), pi.getPlayer().getMap(), 1)) {
               pi.sendPinkNotice("BOSS_ALREADY_STARTED")
               return false
            }
         } else {  //this should never appear
            pi.sendPinkNotice("BOSS_CANNOT_START_YET")
            return false
         }

         pi.playPortalSound()
         return true
      }
   } else {
      pi.playPortalSound()
      pi.warp(922020300, 0)
      return true
   }
}