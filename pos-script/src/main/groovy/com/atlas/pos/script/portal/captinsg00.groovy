package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.hasItem(4000381)) {
      pi.sendPinkNotice("CAPTAIN_LATANICA_MISSING_ESSENCE")
      return false
   } else {
      EventManager em = pi.getEventManager("LatanicaBattle")

      if (pi.getParty().isEmpty()) {
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
   }
}