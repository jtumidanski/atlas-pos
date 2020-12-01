package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(100202)) {
      pi.playPortalSound(); pi.warp(106020400, 2)
      return true
   } else if (pi.hasItem(4000507)) {
      pi.gainItem(4000507, (short) -1)
      pi.sendPinkNotice("POISON_SPORE_USED")

      pi.playPortalSound(); pi.warp(106020400, 2)
      return true
   }

   pi.sendPinkNotice("OVERGROWN_VINES")
   return false
}