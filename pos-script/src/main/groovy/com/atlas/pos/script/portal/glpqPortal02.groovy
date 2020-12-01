package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getJob().getJobNiche() == 2) {
      pi.playPortalSound()
      pi.warp(610030521, 0)
      return true
   } else {
      pi.sendPinkNotice("MAGICIAN_ONLY")
      return false
   }
}