package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(22008)) {
      pi.playPortalSound()
      pi.warp(100030103, "west00")
   } else {
      pi.sendPinkNotice("CANNOT_ENTER_BACKYARD_WITHOUT")
   }
   return true
}