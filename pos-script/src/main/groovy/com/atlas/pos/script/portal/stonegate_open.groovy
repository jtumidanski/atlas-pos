package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("stonegate").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(990000430, 0)
      return true
   } else {
      pi.sendPinkNotice("DOOR_IS_BLOCKED")
      return false
   }
}