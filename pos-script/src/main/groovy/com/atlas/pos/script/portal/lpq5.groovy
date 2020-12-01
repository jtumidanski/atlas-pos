package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int nextMap = 922010700
   EventInstanceManager eim = pi.getPlayer().getEventInstance()
   MapleMap target = eim.getMapInstance(nextMap)
   MaplePortal targetPortal = target.getPortal("st00")
   // only let people through if the eim is ready
   String avail = eim.getProperty("5stageclear")
   if (avail == null) {
      // can't go thru eh?
      pi.sendPinkNotice("SEAL_BLOCKING_DOOR")
      return false
   } else {
      if (eim.getProperty("6stageclear") == null) {
         eim.setProperty("6stageclear", "true")
      }
      pi.playPortalSound()
      pi.getPlayer().changeMap(target, targetPortal)
      return true
   }
}