package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(3366)) {
      pi.playPortalSound()
      pi.warp(926130101, 0)
      return true
   } else {
      pi.sendPinkNotice("ROOM_NO_ACCESS")
      return false
   }
}