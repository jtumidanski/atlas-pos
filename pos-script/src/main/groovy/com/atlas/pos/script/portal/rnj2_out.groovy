package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getEventInstance().getIntProperty("statusStg3") == 3) {
      pi.playPortalSound()
      pi.warp(926100200, 0) //next
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}