package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction
import server.maps.MapleMap
import tools.I18nMessage
import tools.MessageBroadcaster
import tools.ServerNoticeType

boolean enter(PortalPlayerInteraction pi) {
   MapleMap map = pi.getPlayer().getMap()
   if (pi.getPortal().getName() == "female00") {
      if (pi.getPlayer().getGender() == 1) {
         pi.playPortalSound(); pi.warp(map.getId(), "female01")
         return true
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("FEMALE_ONLY"))
         return false
      }
   } else {
      if (pi.getPlayer().getGender() == 0) {
         pi.playPortalSound(); pi.warp(map.getId(), "male01")
         return true
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("MALE_ONLY"))
         return false
      }
   }
}