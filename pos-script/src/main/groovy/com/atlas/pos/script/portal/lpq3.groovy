package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int nextMap = 922010500
   EventInstanceManager eim = pi.getPlayer().getEventInstance()
   MapleMap target = eim.getMapInstance(nextMap)
   MaplePortal targetPortal = target.getPortal("st00")
   // only let people through if the eim is ready
   String avail = eim.getProperty("4stageclear")
   if (avail == null) {
      // can't go thru eh?
      pi.sendPinkNotice("SEAL_BLOCKING_DOOR")
      return false
   } else {
      pi.playPortalSound()
      pi.getPlayer().changeMap(target, targetPortal)
      return true
   }
}