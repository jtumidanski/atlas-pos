package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(1031)) {
      pi.showInfo("UI/tutorial.img/25")
   }

   pi.blockPortal()
   return true
}