package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int baseId = 551030000
   int dungeonId = 551030001
   int dungeons = 19

   if (pi.getMapId() == baseId) {
      if (pi.getParty() != null) {
         if (pi.isLeader()) {
            for (int i = 0; i < dungeons; i++) {
               if (pi.startDungeonInstance(dungeonId + i)) {
                  pi.playPortalSound()
                  pi.warp(dungeonId + i, "out00")
                  return true
               }
            }
         } else {
             pi.sendPinkNotice("PARTY_LEADER_MUST_ENTER")
            return false
         }
      } else {
         for (int i = 0; i < dungeons; i++) {
            if (pi.startDungeonInstance(dungeonId + i)) {
               pi.playPortalSound()
               pi.warp(dungeonId + i, "out00")
               return true
            }
         }
      }
       pi.sendPinkNotice("ALL_MINI_DUNGEON_IN_USE")
      return false
   } else {
      pi.playPortalSound()
      pi.warp(baseId, "MD00")
      return true
   }
}