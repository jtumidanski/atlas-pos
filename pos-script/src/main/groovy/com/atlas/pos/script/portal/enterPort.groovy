package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(21301) && pi.getQuestProgressInt(21301, 9001013) == 0) {
      if (pi.getPlayerCount(108010700) != 0) {
         pi.sendPinkNotice("SOMEONE_ALREADY_CHALLENGING_THIEF_CROW")
         return false
      } else {
         MapleMap map = pi.getClient().getChannelServer().getMapFactory().getMap(108010700)
         spawnMob(2732, 3, 9001013, map)

         pi.playPortalSound()
         pi.warp(108010700, "west00")
      }
   } else {
      pi.playPortalSound(); pi.warp(140020300, 1)
   }
   return true
}

static def spawnMob(x, y, int id, MapleMap map) {
   if (map.getMonsterById(id) != null) {
      return
   }

   MapleLifeFactory.getMonster(id).ifPresent({ mob -> map.spawnMonsterOnGroundBelow(mob, new Point(x, y)) })
}