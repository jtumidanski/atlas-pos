package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

def enter(PortalPlayerInteraction pi) {
   pi.runMapScript()
   return false
}