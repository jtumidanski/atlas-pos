package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("statuegate").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(990000301, 0)
      return true
   } else {
      pi.sendPinkNotice("GATE_IS_CLOSED")
      return false
   }
}