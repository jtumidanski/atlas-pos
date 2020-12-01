package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getJob().getJobNiche() == 4) {
      pi.playPortalSound()
      pi.warp(610030530, 0)
      return true
   } else {
      pi.sendPinkNotice("THIEF_ONLY")
      return false
   }
}