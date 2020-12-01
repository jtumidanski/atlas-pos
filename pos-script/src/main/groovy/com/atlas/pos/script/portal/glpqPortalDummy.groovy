package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   MapleReactor react = pi.getMap().getReactorByName("mob0")

   if (react.getState() < 1) {
      react.forceHitReactor((byte) 1)

      EventInstanceManager eim = pi.getEventInstance()
      eim.setIntProperty("glpq1", 1)

      pi.sendPinkNotice("STRANGE_FORCE")
      pi.playPortalSound(); pi.warp(610030100, 0)

      pi.getEventInstance().showClearEffect()
      eim.giveEventPlayersStageReward(1)
      return true
   }

   pi.sendPinkNotice("PORTAL_MALFUNCTION")
   return false
}