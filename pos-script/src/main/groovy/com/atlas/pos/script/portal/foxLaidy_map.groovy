package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!(pi.isQuestStarted(3647) && pi.hasItem(4031793, 1))) {
      pi.playPortalSound()
      pi.warp(222010200, "east00")
   } else {
      if (!pi.isQuestStarted(23647)) {
         pi.forceStartQuest(23647)
      }
      pi.playPortalSound()
      pi.warp(922220000, "east00")
   }

   return true
}