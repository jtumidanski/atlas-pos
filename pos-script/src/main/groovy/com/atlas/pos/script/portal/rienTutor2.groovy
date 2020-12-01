package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.isQuestCompleted(21011)) {
      pi.sendPinkNotice("COMPLETE_QUEST_BEFORE_PROCEEDING")
      return false
   }
   pi.playPortalSound()
   pi.warp(140090300, 1)
   return true
}