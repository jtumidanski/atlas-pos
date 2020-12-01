package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(20301) || pi.isQuestStarted(20302) || pi.isQuestStarted(20303) || pi.isQuestStarted(20304) || pi.isQuestStarted(20305)) {
      if (pi.hasItem(4032179)) {
         pi.playPortalSound()
         pi.warp(130010000, "east00")
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("NEED_PERMIT"))
         return false
      }
   } else {
      pi.playPortalSound()
      pi.warp(130010000, "east00")
   }
   return true
}