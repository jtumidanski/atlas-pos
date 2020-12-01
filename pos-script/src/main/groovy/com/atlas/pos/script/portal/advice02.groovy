package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInstruction("Press #e#b[Alt]#k#n to\r\\ JUMP.", 100, 5)
   return true
}