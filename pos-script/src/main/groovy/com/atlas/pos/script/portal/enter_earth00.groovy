package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.haveItem(4031890)) {
      pi.sendPinkNotice("WARP_CARD_NEEDED")
      return false
   }

   pi.playPortalSound()
   pi.warp(221000300, "earth00")
   return true
}