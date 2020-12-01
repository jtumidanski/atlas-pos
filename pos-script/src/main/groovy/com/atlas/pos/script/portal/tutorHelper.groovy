package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.spawnGuide()
   pi.talkGuide("Welcome to Maple World! I'm Mimo. I'm in charge of guiding you until you reach Lv. 10 and become a Knight-In-Training. Double-click for further information!")
   pi.blockPortal()
   return true
}