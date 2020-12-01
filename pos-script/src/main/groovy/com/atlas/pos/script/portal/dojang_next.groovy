package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   long currentTime = System.currentTimeMillis()

   if (currentTime - pi.getPlayer().getNpcCoolDown() < 3000) {
      return false
   }
   // this script can be ran twice when passing the dojo portal... strange.
   pi.getPlayer().setNpcCoolDown(currentTime)

   MapleReactor gate = pi.getPlayer().getMap().getReactorByName("door")
   if (gate != null) {
      if (gate.getState() == (byte) 1 || pi.getMap().countMonsters() == 0) {
         if (Math.floor(pi.getPlayer().getMapId() / 100) % 100 < 38) {
            if (((Math.floor((pi.getPlayer().getMap().getId() + 100) / 100)) % 100) % 6 == 0) {
               if (Math.floor(pi.getPlayer().getMapId() / 10000) == 92503) {
                  int restMapId = pi.getPlayer().getMap().getId() + 100
                  int mapId = pi.getPlayer().getMap().getId()

                  for (int i = 0; i < 5; i++) {
                     MapleCharacter[] characters = pi.getMap(mapId - 100 * i).getAllPlayers()

                     Iterator<MapleCharacter> pIter = characters.iterator()
                     while (pIter.hasNext()) {
                        MapleCharacter chr = pIter.next()

                        for (int j = i; j >= 0; j--) {
                           MessageBroadcaster.getInstance().sendServerNotice(chr, ServerNoticeType.PINK_TEXT, I18nMessage.from("DOJO_RECEIVE_POINTS").with(chr.addDojoPointsByMap(mapId - 100 * j), chr.getDojoPoints()))
                        }

                        chr.changeMap(restMapId, 0)
                     }
                  }
               } else {
                  MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("DOJO_RECEIVE_POINTS").with(pi.getPlayer().addDojoPointsByMap(pi.getMapId()), pi.getPlayer().getDojoPoints()))
                  pi.playPortalSound(); pi.warp(pi.getPlayer().getMap().getId() + 100, 0)
               }
            } else {
               MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("DOJO_RECEIVE_POINTS").with(pi.getPlayer().addDojoPointsByMap(pi.getMapId()), pi.getPlayer().getDojoPoints()))
               pi.playPortalSound(); pi.warp(pi.getPlayer().getMap().getId() + 100, 0)
            }
         } else {
            pi.playPortalSound(); pi.warp(925020003, 0)
            pi.getPlayer().gainExp(2000 * pi.getPlayer().getDojoPoints(), true, true, true)
         }
         return true
      } else {
         MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("DOJO_DOOR_NOT_OPEN"))
         return false
      }
   } else {
      return false
   }
}