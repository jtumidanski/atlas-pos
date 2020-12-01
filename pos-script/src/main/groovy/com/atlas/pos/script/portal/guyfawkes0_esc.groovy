package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getEventInstance().getIntProperty("statusStg1") == 1) {
      pi.playPortalSound()
      pi.warp(674030200, 0)
      return true
   } else {
      pi.sendPinkNotice("TUNNEL_BLOCKED")
      return false
   }
}