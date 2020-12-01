package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(6201)) {
      if (pi.getWarpMap(910200000).countPlayers() == 0) {
         pi.resetMapObjects(910200000)
         pi.playPortalSound()
         pi.warp(910200000, 0)

         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_INSIDE")
         return false
      }
   }

   pi.sendPinkNotice("MYSTERIOUS_FORCE")
   return false
}