package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.playPortalSound()
   pi.warp(920010600, 17)
   return true
}