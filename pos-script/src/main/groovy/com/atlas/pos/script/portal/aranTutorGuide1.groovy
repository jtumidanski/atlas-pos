package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.blockPortal()
   if (pi.containsAreaInfo((short) 21002, "chain=o")) {
      return false
   }
   pi.showInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide2")
   pi.sendPinkNotice("ARAN_TUTORIAL_CONSECUTIVE")
   pi.updateAreaInfo((short) 21002, "normal=o;arr0=o;arr1=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
   return true
}