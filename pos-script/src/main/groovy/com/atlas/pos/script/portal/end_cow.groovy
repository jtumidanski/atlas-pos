package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.isQuestStarted(2180) && (pi.hasItem(4031847) || pi.hasItem(4031848) || pi.hasItem(4031849) || pi.hasItem(4031850))) {
      if (pi.hasItem(4031850)) {
         pi.playPortalSound()
         pi.warp(120000103)
         return true
      } else {
         pi.sendPinkNotice("COW_END")
         return false
      }
   } else {
      pi.playPortalSound()
      pi.warp(120000103)
      return true
   }
}