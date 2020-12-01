package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int map = pi.getSavedLocation("MIRROR")
   if (map == -1) {
      map = 100000000
   }

   pi.playPortalSound()
   pi.warp(map)
   return true
}