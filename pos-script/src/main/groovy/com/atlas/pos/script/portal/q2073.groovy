package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(2073)) {
      pi.playPortalSound()
      pi.warp(900000000, 0)
      return true
   } else {
      pi.sendPinkNotice("PRIVATE_PROPERTY")
      return false
   }
}