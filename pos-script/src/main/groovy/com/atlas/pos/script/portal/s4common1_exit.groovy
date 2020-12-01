package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if(pi.hasItem(4031495)) {
      pi.playPortalSound()
      pi.warp(921100301)
   } else {
      pi.playPortalSound()
      pi.warp(211040100)
   }

   return true
}