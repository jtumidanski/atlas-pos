package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getReactorByName("rnj3_out2").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(926100202, 0)
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}