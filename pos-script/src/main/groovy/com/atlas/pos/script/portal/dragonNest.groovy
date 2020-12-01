package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(3706)) {
      pi.playPortalSound()
      pi.warp(240040612, "out00")
      return true
   } else if (pi.isQuestStarted(100203) || pi.getPlayer().haveItem(4001094)) {
      EventManager em = pi.getEventManager("NineSpirit")
      if (!em.startInstance(pi.getPlayer())) {
         pi.sendPinkNotice("SOMEONE_IN_MAP")
         return false
      } else {
         pi.playPortalSound()
         return true
      }
   } else {
      pi.sendPinkNotice("STRANGE_FORCE_2")
      return false
   }
}