package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (!(pi.hasItem(4031507, 5) && pi.hasItem(4031508, 5) && pi.isQuestStarted(6002))) {
      pi.removeAll(4031507)
      pi.removeAll(4031508)
   }

   int pCount = pi.countItem(4031507)
   int rCount = pi.countItem(4031508)

   if (pCount > 5) {
      pi.gainItem(4031507, (short) (-1 * (pCount - 5)))
   }
   if (rCount > 5) {
      pi.gainItem(4031508, (short) (-1 * (rCount - 5)))
   }

   pi.playPortalSound()
   pi.warp(230000003, "out00")
   return true
}