package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int map = pi.getSavedLocation(SavedLocationType.BOSS_PQ.toString())
   if (map == -1) {
      map = 100000000
   }

   pi.playPortalSound()
   pi.warp(map, 0)
   return true
}