package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction


boolean enter(PortalPlayerInteraction pi) {
   pi.playPortalSound()
   pi.warp(pi.getMapId() - 2, 0)
   return true
}