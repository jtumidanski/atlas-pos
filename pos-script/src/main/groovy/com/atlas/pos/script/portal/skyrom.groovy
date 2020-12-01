package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(3935) && !pi.hasItem(4031574, 1)) {
      if(pi.getWarpMap(926000010).countPlayers() == 0) {
         pi.playPortalSound()
         pi.warp(926000010, 0)
         return true
      } else {
         pi.sendPinkNotice("OTHER_PLAYER_TRYING")
         return false
      }
   } else {
      return false
   }
}