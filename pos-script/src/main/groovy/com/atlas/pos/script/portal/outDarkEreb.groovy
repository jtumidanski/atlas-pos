package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int warpMap = pi.isQuestCompleted(20407) ? 924010200 : 924010100

   pi.playPortalSound()
   pi.warp(warpMap, 0)
   return true
}