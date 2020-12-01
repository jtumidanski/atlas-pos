package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInstruction("Press #e#b[Up]#k on the arrow key#n to climb up the ladder or rope.", 350, 5)
   return true
}