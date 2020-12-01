package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasItem(4031582)) {
      pi.playPortalSound()
      pi.warp(260000301, 5)
      return true
   } else {
      pi.sendPinkNotice("ENTRY_PASS_NEEDED")
      return false
   }
}