package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.getEventInstance().isEventCleared()) {
      pi.sendPinkNotice("ZAKUM_COMPLETE_TRIALS")
      return false
   }

   if (pi.getEventInstance().gridCheck(pi.getCharacterId()) == -1) {
      pi.sendPinkNotice("ZAKUM_CLAIM_PRIZE")
      return false
   }

   pi.playPortalSound()
   pi.warp(211042300)
   return true
}