package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasItem(4032125) || pi.hasItem(4032126) || pi.hasItem(4032127) || pi.hasItem(4032128) || pi.hasItem(4032129)) {
      pi.sendPinkNotice("ALREADY_HAVE_PROOF")
      return false
   }

   if (pi.isQuestStarted(20611) || pi.isQuestStarted(20612) || pi.isQuestStarted(20613) || pi.isQuestStarted(20614) || pi.isQuestStarted(20615)) {
      if (pi.getPlayerCount(913020300) == 0) {
         MapleMap map = pi.getMap(913020300)
         map.killAllMonsters()

         pi.playPortalSound()
         pi.warp(913020300, 0)
         pi.spawnMonster(9300294, 87, 88)
         return true
      } else {
         pi.sendPinkNotice("SOMEONE_ALREADY_ATTEMPTING_BOSS")
         return false
      }
   } else {
      pi.sendPinkNotice("CANNOT_ACCESS_HALL")
      return false
   }
}