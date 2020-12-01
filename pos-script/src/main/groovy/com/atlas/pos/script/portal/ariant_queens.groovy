package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isMorphed(221005)) {
      return false
   } else {
      pi.playPortalSound()
      pi.warp(260000300, 7)
      pi.sendPinkNotice("PALACE_INTRUDER")
      return true
   }
}