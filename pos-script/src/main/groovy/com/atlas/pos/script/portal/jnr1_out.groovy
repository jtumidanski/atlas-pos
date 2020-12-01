package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getEventInstance().getIntProperty("statusStg2") == 1) {
      pi.playPortalSound()
      pi.warp(926110100, 0) //next
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}