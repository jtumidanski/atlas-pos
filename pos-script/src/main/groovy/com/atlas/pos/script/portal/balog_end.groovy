package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.canHold(4001261, 1)) {
      pi.sendPinkNotice("INVENTORY_FULL_ERROR")
      return false
   }
   pi.gainItem(4001261, (short) 1)
   pi.playPortalSound()
   pi.warp(105100100, 0)
   return true
}