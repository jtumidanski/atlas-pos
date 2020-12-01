package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().countMonster(9300285) > 0) {
      pi.sendPinkNotice("MUST_DEFEAT_PUPPETEER")
      return false
   } else {
      EventInstanceManager eim = pi.getEventInstance()
      if (eim != null) {
         eim.stopEventTimer()
         eim.dispose()
      }

      pi.playPortalSound()
      pi.warp(105070300, 3)
      return true
   }
}