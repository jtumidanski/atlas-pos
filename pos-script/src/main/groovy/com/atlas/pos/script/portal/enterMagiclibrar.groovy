package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(20718)) {
      EventManager cml = pi.getEventManager("Cygnus_Magic_Library")
      cml.setProperty("player", pi.getPlayer().getName())
      cml.startInstance(pi.getPlayer())
      pi.playPortalSound()
   } else {
      pi.playPortalSound()
      pi.warp(101000003, 8)
   }
   return true
}