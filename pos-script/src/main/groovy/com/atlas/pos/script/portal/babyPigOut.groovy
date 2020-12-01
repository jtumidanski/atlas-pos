package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(22015)) {
      pi.playPortalSound()
      pi.warp(100030300, 2)
   } else {
      pi.sendPinkNotice("RESCUE_BABY_PIG")
   }
   return true
}