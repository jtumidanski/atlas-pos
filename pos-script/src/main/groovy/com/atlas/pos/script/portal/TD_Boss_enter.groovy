package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int stage = ((Math.floor(pi.getMapId() / 100)) % 10) - 1
   EventManager em = pi.getEventManager("TD_Battle" + stage)
   if (em == null) {
      pi.sendPinkNotice("TD_ENCOUNTERED_ERROR", stage)
      return false
   }

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
      } else {
         pi.sendPinkNotice("BOSS_PARTY_MINIMUM")
         return false
      }

      pi.playPortalSound()
      return true
   }
}