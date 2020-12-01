package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.getEventInstance().gridInsert(pi.getPlayer(), 1)
   pi.playPortalSound()
   pi.warp(990000700, "st00")
   return true
}