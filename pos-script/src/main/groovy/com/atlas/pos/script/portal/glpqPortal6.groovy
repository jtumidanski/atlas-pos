package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim != null) {
      if (eim.getIntProperty("glpq6") < 3) {
         pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
         return false
      } else {
         pi.playPortalSound()
         pi.warp(610030700, 0)
         return true
      }
   }

   return false
}