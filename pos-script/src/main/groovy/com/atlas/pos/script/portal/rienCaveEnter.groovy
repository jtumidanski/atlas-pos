package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(21201) || pi.isQuestStarted(21302)) { //aran first job
      pi.playPortalSound()
      pi.warp(140030000, 1)
      return true
   } else {
      pi.sendPinkNotice("SOMETHING_BLOCKING_PORTAL")
      return false
   }
}