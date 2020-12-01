package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInstruction("Press #e#b[Down]#k on the arrow key#n and#e#b[Alt]#k#n at the same time to jump downwards.", 450, 6)
   return true
}