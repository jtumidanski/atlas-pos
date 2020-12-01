package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim != null) {
      if (eim.getIntProperty("glpq2") == 5) {
         pi.playPortalSound()
         pi.warp(610030300, 0)
         return true
      } else {
         pi.sendPinkNotice("PORTAL_NOT_ACTIVE")
         return false
      }
   }

   return false
}