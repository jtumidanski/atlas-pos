package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!pi.haveItem(4031346)) {
      MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("MAGIC_SEED_NEEDED"))
      return false
   }
   if (pi.getMapId() == 240010100) {
      pi.gainItem(4031346, (short) -1)
      pi.playPortalSound()
      pi.warp(101010000, "minar00")
      return true
   } else if (pi.getMapId() == 101010000) {
      pi.gainItem(4031346, (short) -1)
      pi.playPortalSound()
      pi.warp(240010100, "elli00")
      return true
   }
   return true
}