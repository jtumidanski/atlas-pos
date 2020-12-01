package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim != null) {
      if (eim.getIntProperty("glpq5") < 5) {
         pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
         return false
      } else {
         pi.playPortalSound()
         pi.warp(610030600, 0)
         return true
      }
   }

   return false
}