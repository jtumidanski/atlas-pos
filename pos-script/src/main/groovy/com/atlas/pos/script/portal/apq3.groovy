package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.playPortalSound()
   pi.warp(670010600, 0)
   return true
}