package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getReactorByName("jnr31_out").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(926110200, 1)
      return true
   } else {
      pi.sendPinkNotice("DOOR_NOT_YET_OPENED")
      return false
   }
}