package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int returnMap = pi.getSavedLocation("MONSTER_CARNIVAL")
   if (returnMap < 0) {
      returnMap = 102000000 // Just in case there is no saved location.
   }
   MapleMap target = pi.getPlayer().getClient().getChannelServer().getMapFactory().getMap(returnMap)
   pi.getPlayer().changeMap(target)
   pi.playPortalSound()
   return true
}