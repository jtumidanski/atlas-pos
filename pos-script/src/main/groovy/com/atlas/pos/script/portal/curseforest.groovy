package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(2224) || pi.isQuestStarted(2226) || pi.isQuestCompleted(2227)) {
      int hourDay = pi.getHourOfDay()
      if(!((hourDay >= 0 && hourDay < 7) || hourDay >= 17)) {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("CURSED_FOREST_NOT_NOW"))
         return false
      } else {
         pi.playPortalSound(); pi.warp(pi.isQuestCompleted(2227) ? 910100001 : 910100000,"out00")
         return true
      }
   }

   MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("CANNOT_ACCESS"))
   return false
}