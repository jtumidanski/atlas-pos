package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   int evLevel = ((pi.getMapId() - 1) % 5) + 1

   if (pi.getPlayer().getEventInstance().isEventLeader(pi.getPlayer()) && pi.getPlayer().getEventInstance().getPlayerCount() > 1) {
      pi.sendPinkNotice("PARTY_LEADER_CANNOT_LEAVE")
      return false
   }

   if (pi.getPlayer().getEventInstance().giveEventReward(pi.getPlayer(), evLevel)) {
      pi.playPortalSound()
      pi.warp(970030000)
      return true
   } else {
      pi.sendPinkNotice("MAKE_ROOM_AVAILABLE_FOR_PRIZES")
      return false
   }
}