package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMapId() != 777777777) {
      if (!Server.getInstance().canEnterDeveloperRoom()) {
         pi.sendPinkNotice("NEXT_ROOM_NOT_AVAILABLE")
         return false
      }

      pi.saveLocation("DEVELOPER")
      pi.playPortalSound()
      pi.warp(777777777, "out00")
   } else {
      int toMap = pi.getSavedLocation("DEVELOPER")
      pi.playPortalSound()
      pi.warp(toMap, "in00")
   }

   return true
}