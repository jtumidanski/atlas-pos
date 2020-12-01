package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(3928) && pi.isQuestCompleted(3931) && pi.isQuestCompleted(3934)) {
      pi.playPortalSound(); pi.warp(260000201, 1)
      return true
   } else {
      pi.sendPinkNotice("SAND_BANDITS_ONLY")
      return false
   }
}