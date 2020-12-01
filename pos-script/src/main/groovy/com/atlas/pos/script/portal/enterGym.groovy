package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(21701)) {
      pi.playPortalSound()
      pi.warp(914010000, 1)
      return true
   } else if (pi.isQuestStarted(21702)) {
      pi.playPortalSound()
      pi.warp(914010100, 1)
      return true
   } else if (pi.isQuestStarted(21703)) {
      pi.playPortalSound()
      pi.warp(914010200, 1)
      return true
   } else {
      pi.sendPinkNotice("PUO_LESSON_REQUIREMENT")
      return false
   }
}