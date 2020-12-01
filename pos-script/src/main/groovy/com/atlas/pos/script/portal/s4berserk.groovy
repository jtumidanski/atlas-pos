package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(6153) && pi.haveItem(4031475)) {
      MapleMap map = pi.getWarpMap(910500200)
      if (map.countPlayers() == 0) {
         pi.resetMapObjects(910500200)
         map.shuffleReactors()
         pi.playPortalSound(); pi.warp(910500200, "out01")

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