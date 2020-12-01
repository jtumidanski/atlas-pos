package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   String name = pi.getPortal().getName().substring(2, 4)
   MapleReactor gate = pi.getPlayer().getMap().getReactorByName("gate" + name)
   if (gate != null && gate.getState() == (byte) 4) {
      pi.playPortalSound(); pi.warp(670010600, "gt" + name + "PIB")
      return true
   } else {
      pi.sendPinkNotice("GATE_NOT_YET_OPEN")
      return false
   }
}