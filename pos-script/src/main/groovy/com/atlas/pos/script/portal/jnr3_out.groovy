package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getReactorByName("jnr3_out3").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(926110203, 0) //next
      return true
   } else {
      pi.sendPinkNotice("DOOR_NOT_YET_OPENED")
      return false
   }
}