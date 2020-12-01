package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.useItem(2210016)
   pi.playPortalSound()
   pi.warp(200090510, 0)
   return true
}