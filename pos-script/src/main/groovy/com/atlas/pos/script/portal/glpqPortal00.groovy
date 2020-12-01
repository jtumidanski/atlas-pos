package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getJob().getJobNiche() == 1) {
      pi.playPortalSound()
      pi.warp(610030510, 0)
      return true
   } else {
      pi.sendPinkNotice("WARRIOR_ONLY")
      return false
   }
}