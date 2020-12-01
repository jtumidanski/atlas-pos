package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.isQuestStarted(3309) || pi.hasItem(4031708, 1)) {
      pi.playPortalSound()
      pi.warp(261020700, "down00")
   } else {
      pi.playPortalSound()
      pi.warp(926120000, "out00")
   }

   return true
}