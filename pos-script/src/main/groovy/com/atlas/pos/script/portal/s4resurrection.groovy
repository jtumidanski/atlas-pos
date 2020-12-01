package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.hasItem(4001108)) {
      if(pi.getWarpMap(923000100).countPlayers() == 0) {
         pi.resetMapObjects(923000100)
         pi.playPortalSound()
         pi.warp(923000100, 0)

         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_INSIDE")
         return false
      }
   }

   pi.sendPinkNotice("MYSTERIOUS_FORCE")
   return false
}