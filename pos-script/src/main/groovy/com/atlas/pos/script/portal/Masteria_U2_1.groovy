package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasItem(3992040)) {
      pi.playPortalSound()
      pi.warp(610010004, "U5_1")
      return false
   }
   return true
}