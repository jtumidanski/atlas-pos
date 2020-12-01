package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestCompleted(2238)) {
      pi.playPortalSound()
      pi.warp(105100101, "in00")
      return true
   } else {
      pi.sendPinkNotice("MYSTERIOUS_FORCE")
      return false
   }
}