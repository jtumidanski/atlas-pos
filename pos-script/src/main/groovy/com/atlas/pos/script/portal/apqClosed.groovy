package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.sendPinkNotice("GATE_IS_NOT_YET_OPENED")
   return false
}