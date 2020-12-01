package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getReactorByName("rnj3_out1").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(926100201, 0)
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}