package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getReactorByName("jnr32_out").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(926110200, 2)
      return true
   } else {
      pi.sendPinkNotice("DOOR_NOT_YET_OPENED")
      return false
   }
}