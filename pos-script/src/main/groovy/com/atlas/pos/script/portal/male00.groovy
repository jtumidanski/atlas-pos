package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int gender = pi.getPlayer().getGender()
   if (gender == 0) {
      pi.playPortalSound()
      pi.warp(670010200, 3)
      return true
   } else {
      pi.sendPinkNotice("CANNOT_PROCEED")
      return false
   }
}