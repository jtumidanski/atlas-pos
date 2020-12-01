package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.getEventInstance().getIntProperty("statusStg8") == 1) {
      pi.playPortalSound()
      pi.warp(920010920,0)
      return true
   }
   else {
      pi.sendPinkNotice("PIXIE_POWER_REMAINS")
      return false
   }
}