package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(20404)) {
      int warpMap

      if (pi.isQuestCompleted(20407)) {
         warpMap = 924010200
      } else if (pi.isQuestCompleted(20406)) {
         warpMap = 924010100
      } else {
         warpMap = 924010000
      }

      pi.playPortalSound()
      pi.warp(warpMap, 1)
      return true

   } else {
      pi.sendPinkNotice("SHOULD_NOT_GO_CREEPY")
      return false
   }
}