package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   long jailedTime = pi.getJailTimeLeft()

   if (jailedTime <= 0) {
      pi.playPortalSound()
      pi.warp(300000010, "in01")
      return true
   } else {
      int seconds = Math.floor(jailedTime / 1000) % 60
      int minutes = (Math.floor(jailedTime / (1000 * 60)) % 60)
      int hours = (Math.floor(jailedTime / (1000 * 60 * 60)) % 24)

      pi.sendPinkNotice("JAIL_NOTE", hours, minutes, seconds)
      return false
   }
}