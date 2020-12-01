package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.isQuestCompleted(2335) && !(pi.isQuestStarted(2335) && pi.hasItem(4032405))) {
      pi.sendPinkNotice("DOOR_LOCKED")
      return false
   }

   if (pi.isQuestStarted(2335)) {
      pi.forceCompleteQuest(2335, 1300002)
      pi.giveCharacterExp(5000)
      pi.gainItem(4032405, (short) -1)
   }
   pi.playPortalSound()
   pi.warp(106021001, 1)
   return true
}