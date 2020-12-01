package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int mapId = 0

   if (pi.isQuestStarted(20701)) {
      mapId = 913000000
   } else if (pi.isQuestStarted(20702)) {
      mapId = 913000100
   } else if (pi.isQuestStarted(20703)) {
      mapId = 913000200
   }
   if (mapId > 0) {
      if (pi.getPlayerCount(mapId) == 0) {
         MapleMap map = pi.getMap(mapId)
         map.resetPQ()

         pi.playPortalSound()
         pi.warp(mapId, 0)
         return true
      } else {
         pi.sendPinkNotice("SOMEONE_ALREADY_IN_MAP")
         return false
      }
   } else {
      pi.sendPinkNotice("KIKUS_ACCLIMATION_TRAINING_REQUIREMENT")
      return false
   }
}