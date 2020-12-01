package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getReactorByName("kinggate").getState() == (byte) 1) {
      pi.playPortalSound()
      pi.warp(990000900, 2)
      if (pi.getPlayer().getEventInstance().getProperty("boss") != null && pi.getPlayer().getEventInstance().getProperty("boss") == "true") {
         pi.changeMusic("Bgm10/Eregos")
      }
      return true
   } else {
      pi.sendPinkNotice("KING_GATE")
      return false
   }
}