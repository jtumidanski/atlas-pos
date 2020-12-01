package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.isQuestCompleted(2335) && !(pi.isQuestStarted(2335) && pi.hasItem(4032405))) {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("DOOR_LOCKED"))
      return false
   }

   if (pi.isQuestStarted(2335)) {
      pi.forceCompleteQuest(2335, 1300002)
      pi.giveCharacterExp(5000, pi.getPlayer())
      pi.gainItem(4032405, (short) -1)
   }
   pi.playPortalSound()
   pi.warp(106021001, 1)
   return true
}