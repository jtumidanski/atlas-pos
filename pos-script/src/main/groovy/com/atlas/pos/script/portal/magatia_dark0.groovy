package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(7770)) {
      pi.playPortalSound()
      pi.warp(926130000, "out00")
      return true
   } else {
      pi.sendPinkNotice("PIPE_TOO_DARK")
      return false
   }
}