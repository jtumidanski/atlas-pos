package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import server.maps.MapleMap
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   int mapId = 0

   if (pi.isQuestStarted(20701)) {
      mapId = 913000000
   } else if (pi.isQuestStarted(20702)) {
      mapId = 913000100
   } else if (pi.isQuestStarted(20703)) {
      mapId = 913000200
   }
   if (mapId > 0) {
      if (pi.getPlayerCount(mapId) == 0) {
         MapleMap map = pi.getMap(mapId)
         map.resetPQ()

         pi.playPortalSound()
         pi.warp(mapId, 0)
         return true
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("SOMEONE_ALREADY_IN_MAP"))
         return false
      }
   } else {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("KIKUS_ACCLIMATION_TRAINING_REQUIREMENT"))
      return false
   }
}