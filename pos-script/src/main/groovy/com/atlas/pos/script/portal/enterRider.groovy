package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(21610) && pi.haveItem(4001193, 1)) {
      EventManager em = pi.getEventManager("Aran_2ndmount")
      if (em == null) {
         pi.sendPinkNotice("2ND_MOUNT_QUEST_CLOSED")
         return false
      } else {
         if (!em.startInstance(pi.getPlayer())) {
            pi.sendPinkNotice("2ND_MOUNT_QUEST_SOMEONE_ALREADY_IN")
            return false
         } else {
            pi.playPortalSound()
            return true
         }
      }
   } else {
      pi.sendPinkNotice("2ND_MOUNT_QUEST_REQUIREMENT")
      return false
   }
}