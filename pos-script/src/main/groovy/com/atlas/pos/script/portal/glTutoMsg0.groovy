package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInstruction("Once you leave this area you won't be able to return.", 150, 5)
   return true
}