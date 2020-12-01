package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(22010) || pi.getJobId() != 2001) {
      pi.playPortalSound()
      pi.warp(100030310, 0)
   } else {
      pi.sendPinkNotice("CANNOT_ENTER_LUSH_FOREST_WITHOUT")
   }
   return true
}