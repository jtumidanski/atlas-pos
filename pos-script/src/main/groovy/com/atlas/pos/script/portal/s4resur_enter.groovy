package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(6134)) {
      pi.playPortalSound()
      pi.warp(922020000, 0)
      return true
   }

   pi.sendPinkNotice("MYSTERIOUS_FORCE")
   return false
}