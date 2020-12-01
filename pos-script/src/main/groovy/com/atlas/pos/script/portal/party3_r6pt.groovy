package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim.getProperty("stage6_comb") == null) {
      String comb = "0"

      for (int i = 0; i < 16; i++) {
         int r = Math.floor((Math.random() * 4)).intValue() + 1
         comb += r.toString()
      }

      eim.setProperty("stage6_comb", comb)
   }

   String comb = eim.getProperty("stage6_comb")

   String name = pi.getPortal().name().substring(2, 5)
   int portalId = Integer.parseInt(name)


   int pRow = (int) Math.floor(portalId / 10)
   int pCol = portalId % 10

   if (pCol == Integer.parseInt(comb.substring(pRow, pRow + 1))) {    //climb
      pi.playPortalSound()
      pi.warp(pi.getMapId(), (pRow % 4 != 0) ? pi.getPortal().id() + 4 : ((int) (pRow / 4)))
   } else {    //fail
      pRow--
      pi.playPortalSound()
      pi.warp(pi.getMapId(), (pRow / 4) > 1 ? ((int) (pRow / 4)) : 5)
   }

   return true
}