package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(6230) || pi.isQuestStarted(6231) || pi.haveItem(4001110)) {
      if(pi.getWarpMap(922020200).countPlayers() == 0) {
         pi.resetMapObjects(922020200)
         pi.playPortalSound()
         pi.warp(922020200, 0)

         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_INSIDE")
         return false
      }
   }

   return false
}