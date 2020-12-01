package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventInstanceManager eim = pi.getEventInstance()
   int area = eim.getIntProperty("statusStg5")
   int reg = 2

   if((area >> reg) % 2 == 0) {
      area |= (1 << reg)
      eim.setIntProperty("statusStg5", area)

      pi.playPortalSound(); pi.warp(926110301 + reg, 0) //next
      return true
   } else {
      pi.sendPinkNotice("ROOM_ALREADY_BEING_EXPLORED")
      return false
   }
}