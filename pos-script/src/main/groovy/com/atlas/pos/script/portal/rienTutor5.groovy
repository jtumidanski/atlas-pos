package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.talkGuide("You're very close to town. I'll head over there first since I have some things to take care of. You take your time.")
   pi.blockPortal()
   return false
}