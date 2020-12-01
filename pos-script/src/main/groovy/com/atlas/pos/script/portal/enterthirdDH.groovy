package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.hasItem(4032120) || pi.hasItem(4032121) || pi.hasItem(4032122) || pi.hasItem(4032123) || pi.hasItem(4032124)) {
      pi.sendPinkNotice("ALREADY_HAVE_PROOF_OF_QUALIFICATION")
      return false
   }
   if (pi.isQuestStarted(20601) || pi.isQuestStarted(20602) || pi.isQuestStarted(20603) || pi.isQuestStarted(20604) || pi.isQuestStarted(20605)) {
      if (pi.getPlayerCount(913010200) == 0) {
         MapleMap map = pi.getMap(913010200)
         map.killAllMonsters()
         pi.playPortalSound(); pi.warp(913010200, 0)
         pi.spawnMonster(9300289, 0, 0)
         return true
      } else {
         pi.sendPinkNotice("SOMEONE_ALREADY_ATTEMPTING_BOSS")
         return false
      }
   } else {
      pi.sendPinkNotice("LEVEL_100_SKILL_REQUIREMENT")
      return false
   }
}