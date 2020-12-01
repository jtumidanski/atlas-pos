package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isEventLeader()) {
      EventInstanceManager eim = pi.getPlayer().getEventInstance()
      int target
      byte theWay = pi.getMap().getReactorByName("light").getState()
      if (theWay == (byte) 1) {
         target = 240050300 //light
      } else if (theWay == (byte) 3) {
         target = 240050310 //dark
      } else {
         pi.sendPinkNotice("HORNTAIL_HIT_LIGHT_BULB")
         return false
      }

      pi.playPortalSound()
      eim.warpEventTeam(target)
      return true
   } else {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("HORNTAIL_PARTY_LEADER"))
      return false
   }
}