package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.unlockUI()
   return true
}