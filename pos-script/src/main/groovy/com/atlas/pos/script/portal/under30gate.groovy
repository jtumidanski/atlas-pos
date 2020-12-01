package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getLevel() <= 30) {
      pi.playPortalSound()
      pi.warp(990000640, 1)
      return true
   } else {
      pi.sendPinkNotice("CANNOT_PROCEED_PAST_POINT")
      return false
   }
}