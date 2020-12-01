package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("secretgate2").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(990000631,1)
      return true
   } else {
      pi.sendPinkNotice("DOOR_IS_CLOSED")
      return false
   }
}