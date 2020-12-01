package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int[] backPortals = [6, 8, 9, 11]
   int idx = pi.getEventInstance().gridCheck(pi.getPlayer())

   pi.playPortalSound(); pi.warp(990000600, backPortals[idx])
   return true
}