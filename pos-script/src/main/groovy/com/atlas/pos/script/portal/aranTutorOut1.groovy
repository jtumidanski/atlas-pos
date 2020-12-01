package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(21000)) {
      //lol nexon does this xD
      pi.teachSkill(20000017, (byte) 0, (byte) -1, -1)
      pi.teachSkill(20000018, (byte) 0, (byte) -1, -1)
      //nexon sends updatePlayerStats Ma(byte) 1, (byte) 0, -1);
      pi.teachSkill(20000018, (byte) 1, (byte) 0, -1)
      //actually nexon does enableActions here :P
      pi.playPortalSound()
      pi.warp(914000200, 1)
      return true
   } else {
      pi.sendPinkNotice("ARAN_TUTORIAL_EXIT")
      return false
   }
}