package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getPlayer().getEventInstance()
   MapleMap target = eim.getMapInstance(922010800)
   if (eim.getProperty("7stageclear") != null) {
      pi.playPortalSound()
      pi.getPlayer().changeMap(target, target.getPortal("st00"))
      return true
   } else {
      return false
   }
}