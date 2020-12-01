package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int mapId = pi.getSavedLocation("MIRROR")

   pi.playPortalSound()
   if (mapId == -1) {
      pi.warp(102040000, 12)
   } else {
      pi.warp(mapId)
   }

   //pi.warp(102040000, 12);
   return true
}