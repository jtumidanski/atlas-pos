package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int toMap = pi.getSavedLocation("FREE_MARKET")
   pi.playPortalSound()
   pi.warp(toMap, pi.getMarketPortalId(toMap))
   return true
}