package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(6241) || pi.isQuestStarted(6243)) {
      if(pi.getWarpMap(924000100).countPlayers() == 0) {
         pi.resetMapObjects(924000100)
         pi.playPortalSound()
         pi.warp(924000100, 0)

         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_INSIDE")
         return false
      }
   }

   pi.sendPinkNotice("MYSTERIOUS_FORCE")
   return false
}