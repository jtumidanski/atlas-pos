package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   switch (pi.getMapId()) {
      case 930000000:
         pi.playPortalSound()
         pi.warp(930000100, 0)
         return true
         break
      case 930000100:
         if (pi.getMap().getMonsters().size() == 0) {
            pi.playPortalSound()
            pi.warp(930000200, 0)
            return true
         } else {
            pi.sendPinkNotice("ELIMINATE_MONSTERS")
            return false
         }
         break
      case 930000200:
         if (pi.getMap().getReactorByName("spine") != null && pi.getMap().getReactorByName("spine").getState() < 4) {
            pi.sendPinkNotice("SPINE_BLOCKS")
            return false
         } else {
            pi.playPortalSound()
            pi.warp(930000300, 0) //assuming they cant get past reactor without it being gone
            return true
         }
         break

      default:
         pi.sendPinkNotice("UNBOUND_PATH")
         return false
   }
}