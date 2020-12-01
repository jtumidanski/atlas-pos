package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getPlayer().getMap().getMonsterById(9300216) != null) {
      pi.getPlayer().enteredScript("dojang_Msg", pi.getPlayer().getMap().getId())
      pi.getPlayer().setFinishedDojoTutorial()
      pi.getClient().getChannelServer().resetDojo(pi.getPlayer().getMap().getId())
      pi.getClient().getChannelServer().dismissDojoSchedule(pi.getPlayer().getMap().getId(), pi.getParty().get())
      pi.playPortalSound(); pi.warp(925020001, 0)
      return true
   } else {
      pi.sendPinkNotice("DOJO_DONT_RUN")
      return false
   }
}