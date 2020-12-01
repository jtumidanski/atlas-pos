package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().countMonsters() == 0) {
      if (pi.canHold(4001193, 1)) {
         pi.gainItem(4001193, (short) 1)
         pi.playPortalSound()
         pi.warp(140010210, 0)
         return true
      } else {
         pi.sendPinkNotice("FREE_SPACE_FOR_COUSE_CLEAR_TOKEN")
         return false
      }
   } else {
      pi.sendPinkNotice("DEFEAT_ALL_WOLVES")
      return false
   }
}