package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int portal = 0
   switch (pi.getTeam()) {
      case 0:
         portal = 4
         break
      case 1:
         portal = 3
         break
   }
   pi.warp(980000301, portal)
   pi.playPortalSound()
   return true
}