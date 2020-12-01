package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(2324)) {
      pi.playPortalSound()
      pi.warp(106020501, 0)
      return true
   } else {
      pi.sendPinkNotice("NEED_THORN_REMOVER")
      return false
   }
}