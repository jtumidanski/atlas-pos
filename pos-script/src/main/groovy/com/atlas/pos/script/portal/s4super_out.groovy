package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int exit = pi.getEventInstance().getIntProperty("canLeave")
   if (exit == 0) {
      pi.sendPinkNotice("WAIT_TO_LEAVE")
      return false
   } else if (exit == 2) {
      pi.playPortalSound()
      pi.warp(912010200)
      return true
   } else {
      pi.playPortalSound()
      pi.warp(120000101)
      return true
   }
}