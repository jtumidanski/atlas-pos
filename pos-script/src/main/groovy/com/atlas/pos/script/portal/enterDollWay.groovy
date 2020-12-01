package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestCompleted(20730) || pi.isQuestCompleted(21734)) {  // puppeteer defeated, new found secret path
      pi.playPortalSound(); pi.warp(105070300, 3)
      return true
   } else if (pi.isQuestStarted(21734)) {
      pi.playPortalSound(); pi.warp(910510100, 0)
      return true
   } else {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("OMINOUS_POWER"))
      return false
   }
}