package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("ghostgate").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(990000800, 0)
      return true
   } else {
      pi.sendPinkNotice("NOT_OPEN_YET")
      return false
   }
}