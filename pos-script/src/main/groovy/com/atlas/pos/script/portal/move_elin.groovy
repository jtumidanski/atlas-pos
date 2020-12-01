package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.playPortalSound()
   pi.warp(300000100, "out00")
   pi.sendPinkNotice("PASSING_TIME_GATE")
   return true
}