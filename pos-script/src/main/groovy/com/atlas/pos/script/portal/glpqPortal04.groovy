package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getJob().getJobNiche() == 5) {
      pi.playPortalSound()
      pi.warp(610030550, 0)
      return true
   } else {
      pi.sendPinkNotice("PIRATE_ONLY")
      return false
   }
}