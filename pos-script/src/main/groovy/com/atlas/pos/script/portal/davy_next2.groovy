package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().getMonsters().size() == 0 && passedGrindMode(pi.getMap(), pi.getEventInstance())) {
      pi.playPortalSound()
      pi.warp(925100300, 0) //next
      return true
   } else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}

static def passedGrindMode(MapleMap map, EventInstanceManager eim) {
   if (eim.getIntProperty("grindMode") == 0) {
      return true
   }
   return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
}