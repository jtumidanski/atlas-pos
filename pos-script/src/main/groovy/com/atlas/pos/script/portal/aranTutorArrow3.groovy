package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.blockPortal()
   if (pi.containsAreaInfo((short) 21002, "arr3=o")) {
      return false
   }
   pi.updateAreaInfo((short) 21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
   pi.showInfo("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
   return true
}