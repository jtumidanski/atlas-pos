package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(1008)) {
      pi.showInfo("UI/tutorial.img/22")
   } else if (pi.isQuestCompleted(1020)) {
      pi.showInfo("UI/tutorial.img/27")
   }

   pi.blockPortal()
   return true
}