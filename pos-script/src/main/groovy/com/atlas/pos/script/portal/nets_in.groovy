package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.saveLocation("MIRROR")
   pi.playPortalSound()
   pi.warp(926010000, 4)
   return true
}