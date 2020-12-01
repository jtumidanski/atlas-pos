package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   if (eim != null) {
      eim.stopEventTimer()
      eim.dispose()
   }

   //3 Yetis
   int questProgress = pi.getQuestProgressInt(2330, 3300005) + pi.getQuestProgressInt(2330, 3300006) + pi.getQuestProgressInt(2330, 3300007)
   if (questProgress == 3 && !pi.hasItem(4032388)) {
      if (pi.canHold(4032388)) {
         pi.sendPinkNotice("PEPE_KING_DROP")
         pi.gainItem(4032388, (short) 1)

         pi.playPortalSound()
         pi.warp(106021400, 2)
         return true
      } else {
         pi.sendPinkNotice("INVENTORY_FULL", "ETC")
         return false
      }
   } else {
      pi.playPortalSound()
      pi.warp(106021400, 2)
      return true
   }
}