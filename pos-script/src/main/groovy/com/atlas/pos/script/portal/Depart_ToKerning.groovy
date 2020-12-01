package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventManager em = pi.getEventManager("KerningTrain")
   if (!em.startInstance(pi.getPlayer())) {
      pi.sendPinkNotice("WAGON_FULL")
      return false
   }

   pi.playPortalSound()
   return true
}