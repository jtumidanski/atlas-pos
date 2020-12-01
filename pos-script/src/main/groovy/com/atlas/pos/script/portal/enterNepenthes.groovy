package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestActive(21739)) {
      MapleMap map1 = pi.getWarpMap(920030000)
      MapleMap map2 = pi.getWarpMap(920030001)

      if (map1.countPlayers() == 0 && map2.countPlayers() == 0) {
         map1.resetPQ(1)
         map2.resetPQ(1)

         MapleLifeFactory.getMonster(9300348).ifPresent({ monster -> map2.spawnMonsterOnGroundBelow(monster, new Point(591, -34)) })

         pi.playPortalSound(); pi.warp(920030000, 2)
         return true
      } else {
         pi.sendPinkNotice("SOMEONE_ALREADY_CHALLENGING")
         return false
      }
   } else {
      pi.playPortalSound(); pi.warp(200060001, 2)
      return true
   }
}