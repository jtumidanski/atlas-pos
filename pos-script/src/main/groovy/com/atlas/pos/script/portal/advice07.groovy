package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInstruction("You can view the World Map by pressing the #e#b[W]#k#nkey.", 350, 5)
   return true
}