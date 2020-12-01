package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getJob().getJobNiche() == 3) {
      pi.playPortalSound()
      pi.warp(610030540, 0)
      return true
   } else {
      pi.sendPinkNotice("ARCHER_ONLY")
      return false
   }
}