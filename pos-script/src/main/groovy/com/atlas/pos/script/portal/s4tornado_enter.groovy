package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   if(pi.isQuestStarted(6230) || pi.isQuestStarted(6231) || pi.haveItem(4001110)) {
      if(pi.getWarpMap(922020200).countPlayers() == 0) {
         pi.resetMapObjects(922020200)
         pi.playPortalSound()
         pi.warp(922020200, 0)

         return true
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("OTHER_PLAYER_INSIDE"))
         return false
      }
   }

   return false
}