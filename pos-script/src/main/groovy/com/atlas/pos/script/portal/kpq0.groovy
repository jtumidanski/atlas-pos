package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getPlayer().getEventInstance()
   MapleMap target = eim.getMapInstance(103000801)

   if (eim.getProperty("1stageclear") != null) {
      pi.playPortalSound()
      pi.getPlayer().changeMap(target, target.getPortal("st00"))
      return true
   }
   else {
      pi.sendPinkNotice("PORTAL_NOT_YET_OPENED")
      return false
   }
}