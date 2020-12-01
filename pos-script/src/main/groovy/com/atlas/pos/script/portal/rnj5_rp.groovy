package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   String mapPlayer = "stage6_comb" + (pi.getMapId() % 10)
   EventInstanceManager eim = pi.getEventInstance()

   String comb = eim.getProperty(mapPlayer)

   String name = pi.getPortal().name().substring(2, 4)
   int portalId = Integer.parseInt(name)


   int pRow = (int) Math.floor(portalId / 10)
   int pCol = (portalId % 10)

   if (pCol == Integer.parseInt(comb.substring(pRow, pRow + 1))) {    //climb
      if (pRow < 9) {
         pi.playPortalSound(); pi.warp(pi.getMapId(), pi.getPortal().id() + 4)
      } else {
         if (eim.getIntProperty("statusStg6") == 0) {
            eim.setIntProperty("statusStg6", 1)
            eim.giveEventPlayersStageReward(6)
         }

         pi.playPortalSound()
         pi.warp(pi.getMapId(), 1)
      }

   } else {    //fail
      pi.playPortalSound()
      pi.warp(pi.getMapId(), 2)
   }

   return true
}