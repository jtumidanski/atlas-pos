package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.blockPortal()
   if (pi.containsAreaInfo((short) 21002, "normal=o")) {
      return false
   }
   pi.showInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide1")
   pi.sendPinkNotice("ARAN_TUTORIAL_REGULAR_ATTACK")
   pi.updateAreaInfo((short) 21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o")
   return true
}