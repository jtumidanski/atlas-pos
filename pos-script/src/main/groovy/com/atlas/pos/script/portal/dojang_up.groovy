package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getMonsterById(9300216) != null) {
      pi.goDojoUp()
      pi.getPlayer().getMap().setReactorState()
      int stage = Math.floor(pi.getPlayer().getMapId() / 100) % 100
      if ((stage - (stage / 6) | 0) == pi.getPlayer().getVanquisherStage() && !GameConstants.isDojoPartyArea(pi.getPlayer().getMapId())) // we can also try 5 * stage / 6 | 0 + 1
      {
         pi.getPlayer().setVanquisherKills(pi.getPlayer().getVanquisherKills() + 1)
      }
   } else {
      pi.sendPinkNotice("DOJO_MORE_MONSTERS")
   }
   pi.enableActions()
   return true
}