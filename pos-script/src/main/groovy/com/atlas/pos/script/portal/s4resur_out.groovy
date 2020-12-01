package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(6134)) {
      if(pi.canHold(4031448)) {
         pi.gainItem(4031448, (short) 1)
         pi.playPortalSound()
         pi.warp(220070400, 3)

         return true
      } else {
         pi.sendPinkNotice("MAKE_ROOM_FOR_QUEST_ITEM", "ETC")
         return false
      }
   } else {
      pi.playPortalSound()
      pi.warp(220070400, 3)
      return true
   }
}