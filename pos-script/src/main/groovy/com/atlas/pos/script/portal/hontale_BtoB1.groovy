package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getMap().countPlayers() == 1) {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("HORNTAIL_LAST_PLAYER"))
      return false
   } else {
      if (pi.haveItem(4001087)) {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("HORNTAIL_CANNOT_PASS_WITH_1ST_CRYSTAL"))
         return false
      }
      pi.playPortalSound()
      pi.warp(240050101, 0)
      return true
   }
}