package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   MapleMap map = pi.getPlayer().getMap()
   if (pi.getPortal().getName() == "female00") {
      if (pi.getGender() == 1) {
         pi.playPortalSound()
         pi.warp(map.getId(), "female01")
         return true
      } else {
         pi.sendPinkNotice("FEMALE_ONLY")
         return false
      }
   } else {
      if (pi.getGender() == 0) {
         pi.playPortalSound()
         pi.warp(map.getId(), "male01")
         return true
      } else {
         pi.sendPinkNotice("MALE_ONLY")
         return false
      }
   }
}