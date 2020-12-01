package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasLevel30Character()) {
      pi.openNpc(2007)
   }
   pi.blockPortal()
   return true
}