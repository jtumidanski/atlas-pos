package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().countMonsters() == 0) {
      pi.playPortalSound()
      pi.warp(910500200, "out00")
      return true
   }
   pi.sendPinkNotice("DEFEAT_ALL_MONSTERS_FIRST")
   return true
}