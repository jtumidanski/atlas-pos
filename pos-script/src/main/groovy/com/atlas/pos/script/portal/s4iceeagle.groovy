package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(6242)) {
      if (pi.getWarpMap(921100210).countPlayers() == 0) {
         pi.resetMapObjects(921100210)
         pi.playPortalSound()
         pi.warp(921100210, 0)

         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_INSIDE")
         return false
      }
   } else {
      pi.sendPinkNotice("MYSTERIOUS_FORCE")
      return false
   }
}