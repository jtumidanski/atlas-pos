package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

def enter(PortalPlayerInteraction pi) {
   pi.runMapScript()
   return false
}