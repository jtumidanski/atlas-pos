package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   long entryTime = pi.getPlayer().getEventInstance().getProperty("entryTimestamp").toLong()
   long timeNow = System.currentTimeMillis()

   int timeLeft = Math.ceil((entryTime - timeNow) / 1000).toInteger()

   if(timeLeft <= 0) {
      pi.playPortalSound(); pi.warp(990000100, 0)
      return true
   }
   else { //cannot proceed while allies can still enter
      pi.sendPinkNotice("PORTAL_OPEN_IN", timeLeft)
      return false
   }
}