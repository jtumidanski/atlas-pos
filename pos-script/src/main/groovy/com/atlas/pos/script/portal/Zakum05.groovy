package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!(pi.isQuestStarted(100200) || pi.isQuestCompleted(100200))) {
      pi.sendPinkNotice("ZAKUM_MASTER_APPROVAL")
      return false
   }

   if (!pi.isQuestCompleted(100201)) {
      pi.sendPinkNotice("ZAKUM_COMPLETE_TRIALS_2")
      return false
   }

   if (!pi.haveItem(4001017)) {
      pi.sendPinkNotice("ZAKUM_NEED_EYE_OF_FIRE")
      return false
   }

   MapleReactor react = pi.getMap().getReactorById(2118002)
   if (react != null && react.getState() > 0) {
      pi.sendPinkNotice("ENTRANCE_BLOCKED")
      return false
   }

   pi.playPortalSound()
   pi.warp(211042400, "west00")
   return true
}