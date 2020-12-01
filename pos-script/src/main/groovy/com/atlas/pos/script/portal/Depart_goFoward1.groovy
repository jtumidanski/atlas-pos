package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int mapId = pi.getMapId()

   if (mapId == 103040410 && pi.isQuestCompleted(2287)) {
      pi.playPortalSound()
      pi.warp(103040420, "right01")
      return true
   } else if (mapId == 103040420 && pi.isQuestCompleted(2288)) {
      pi.playPortalSound()
      pi.warp(103040430, "right01")
      return true
   } else if (mapId == 103040410 && pi.isQuestStarted(2287)) {
      pi.playPortalSound()
      pi.warp(103040420, "right01")
      return true
   } else if (mapId == 103040420 && pi.isQuestStarted(2288)) {
      pi.playPortalSound()
      pi.warp(103040430, "right01")
      return true
   } else {
      if (mapId == 103040440 || mapId == 103040450) {
         pi.playPortalSound()
         pi.warp(mapId + 10, "right01")
         return true
      }
      pi.sendPinkNotice("CANNOT_ACCESS")
      return false
   }
}