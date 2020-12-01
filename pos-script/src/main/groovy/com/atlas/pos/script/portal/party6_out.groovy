package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()

   if (eim.isEventCleared()) {
      if(pi.isEventLeader()) {
         pi.playPortalSound()
         eim.warpEventTeam(930000800)
         return true
      } else {
         pi.sendPinkNotice("WAIT_FOR_LEADER")
         return false
      }
   } else {
      pi.sendPinkNotice("ELIMINATE_POISON_GOLEM")
      return false
   }
}