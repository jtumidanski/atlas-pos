package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(1035)) {
      pi.showInfo("UI/tutorial.img/21")
   }

   pi.blockPortal()
   return true
}