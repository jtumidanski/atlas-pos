package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(1035)) {
      pi.showInfo("UI/tutorial.img/23")
   }

   pi.blockPortal()
   return true
}