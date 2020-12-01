package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.blockPortal()
   if (pi.containsAreaInfo((short) 21002, "cmd=o")) {
      return false
   }
   pi.showInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide3")
   pi.sendPinkNotice("ARAN_TUTORIAL_COMMAND")
   pi.updateAreaInfo((short) 21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
   return true
}