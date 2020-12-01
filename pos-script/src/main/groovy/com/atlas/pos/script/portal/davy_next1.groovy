package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim != null && eim.getProperty("stage2") == "3") {
      pi.playPortalSound()
      pi.warp(925100200, 0) //next
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
   return false
}