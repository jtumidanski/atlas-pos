package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getPlayer().getEventInstance()
   int fc = eim.getIntProperty("falls")

   if (fc >= 3) {
      pi.playPortalSound()
      pi.warp(922240200, 0)
   } else {
      eim.setIntProperty("falls", fc + 1)
      pi.playPortalSound()
      pi.warp(pi.getMapId(), 0)
   }

   return true
}