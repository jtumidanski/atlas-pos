package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getEventInstance().getIntProperty("glpq1") == 0) {
      pi.sendPinkNotice("PATH_BLOCKED")
      return false

   } else {
      pi.playPortalSound()
      pi.warp(610030200, 0)
      return true
   }
}