package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestActive(21747) || pi.isQuestActive(21744) && pi.isQuestCompleted(21745)) {
      pi.playPortalSound()
      pi.warp(925040000, 0)
      return true
   } else {
      pi.sendPinkNotice("NO_PERMISSION_TO_ENTER")
      return false
   }
}