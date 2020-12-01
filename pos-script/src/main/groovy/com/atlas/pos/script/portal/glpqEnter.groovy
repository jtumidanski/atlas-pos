package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasItem(3992041, 1)) {
      pi.playPortalSound()
      pi.warp(610030020, "out00")
      return true
   } else {
      pi.sendPinkNotice("GIANT_GATE_NO_BUDGE")
      return false
   }
}