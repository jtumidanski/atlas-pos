package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   //TODO
   // pi.setDirectionStatus(true)
   pi.lockUI()
   pi.openNpc(3, "1096003")
   return true
}