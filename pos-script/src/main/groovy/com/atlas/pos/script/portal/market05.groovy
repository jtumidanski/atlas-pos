package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMapId() != 910000000) {
      pi.saveLocation("FREE_MARKET")
      pi.playPortalSound()
      pi.warp(910000000, "out00")
      return true
   }
   return false
}