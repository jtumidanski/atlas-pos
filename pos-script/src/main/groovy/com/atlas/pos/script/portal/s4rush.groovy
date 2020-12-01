package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(6110)) {
      pi.playPortalSound()
      pi.warp(910500100, 0)
      return true
   } else {
      pi.sendPinkNotice("MYSTERIOUS_FORCE")
      return false
   }
}