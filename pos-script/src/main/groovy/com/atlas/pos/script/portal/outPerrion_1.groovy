package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.sendPinkNotice("TEMPLE_SHORTCUT")
   pi.playPortalSound()
   pi.warp(105100000, 2)
   return true
}