package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int mapId = pi.getSavedLocation("MIRROR")
   pi.playPortalSound()
   if(mapId == 260020500) pi.warp(mapId, 3)
   else pi.warp(mapId)
   return true
}