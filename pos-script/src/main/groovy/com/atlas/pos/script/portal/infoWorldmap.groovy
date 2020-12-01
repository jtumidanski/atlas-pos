package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.showInfo("UI/tutorial.img/26")
   pi.blockPortal()
   return true
}