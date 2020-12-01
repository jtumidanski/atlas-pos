package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(1041)) {
      pi.playPortalSound()
      pi.warp(1010100, 4)
   } else if (pi.isQuestStarted(1042)) {
      pi.playPortalSound()
      pi.warp(1010200, 4)
   } else if (pi.isQuestStarted(1043)) {
      pi.playPortalSound()
      pi.warp(1010300, 4)
   } else if (pi.isQuestStarted(1044)) {
      pi.playPortalSound()
      pi.warp(1010400, 4)
   } else {
      pi.sendPinkNotice("MAI_TRAINING_REQUIREMENT")
      return false
   }
   return true
}