package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int jobType = 3

   if (pi.isQuestStarted(20301) || pi.isQuestStarted(20302) || pi.isQuestStarted(20303) || pi.isQuestStarted(20304) || pi.isQuestStarted(20305)) {
      MapleMap map = pi.getClient().getChannelServer().getMapFactory().getMap(108010600 + (10 * jobType))
      if (map.countPlayers() > 0) {
         pi.sendPinkNotice("SOMEONE_ALREADY_SEARCHING")
         return false
      }

      if (pi.hasItem(4032101 + jobType, 1)) {
         pi.sendPinkNotice("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
         return false
      }

      pi.playPortalSound(); pi.warp(108010600 + (10 * jobType), "out00")
   } else {
      pi.playPortalSound(); pi.warp(130010120, "out00")
   }
   return true
}