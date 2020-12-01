package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.getEventInstance().isEventCleared()) {
      pi.sendPinkNotice("COMPLETE_MISSION_BEFORE_PROCEEDING")
      return false
   } else {
      if (pi.isQuestStarted(6410)) {
         pi.setQuestProgress(6410, 6411, "p2")
      }

      pi.playPortalSound()
      pi.warp(925010400)
      return true
   }
}