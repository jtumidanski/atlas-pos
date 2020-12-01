package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().countMonster(2220100) > 0) {
      pi.sendPinkNotice("DEFEAT_BLUE_MUSHROOMS")
      return false
   } else {
      EventInstanceManager eim = pi.getEventInstance()
      eim.stopEventTimer()
      eim.dispose()

      pi.playPortalSound()
      pi.warp(101000000, 26)

      if (pi.isQuestCompleted(20718)) {
         pi.openNpc(1103003, "MaybeItsGrendel_end")
      }

      return true
   }
}