package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("speargate").getState() == (byte) 4) {
      pi.playPortalSound()
      pi.warp(990000401, 0)
      return true
   } else {
      pi.sendPinkNotice("NOT_OPEN_YET")
      return false
   }
}