package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.sendPinkNotice("SEEMS_TO_BE_LOCKED")
   return true
}