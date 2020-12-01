package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getJobId() == 2000 && !pi.isQuestCompleted(21014)) {
      pi.playPortalSound()
      pi.warp(140000000, "st00")
   } else {
      pi.playPortalSound()
      pi.warp(140000000, "west00")
   }

   return true
}