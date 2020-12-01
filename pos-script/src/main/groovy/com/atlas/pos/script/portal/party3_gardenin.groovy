package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getParty().isPresent() && pi.isEventLeader() && pi.hasItem(4001055, 1)) {
      pi.playPortalSound()
      pi.getEventInstance().warpEventTeam(920010100)
      return true
   } else {
      pi.sendPinkNotice("NEED_TO_BE_LEADER_AND_HAVE_ROOT_OF_LIFE")
      return false
   }
}