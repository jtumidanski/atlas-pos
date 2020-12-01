package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int gender = pi.getGender()
   if (gender == 1) {
      pi.playPortalSound(); pi.warp(670010200, 4)
      return true
   } else {
      pi.sendPinkNotice("CANNOT_PROCEED")
      return false
   }
}