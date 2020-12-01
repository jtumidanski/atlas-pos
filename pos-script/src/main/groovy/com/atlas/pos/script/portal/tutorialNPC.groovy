package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   if (pi.getLevel() <= 10 && pi.getJobId() == 0) {
      int m = pi.getMapId()
      int npcId = 0

      if (m == 120000101) { // Navigation Room, The Nautilus
         npcId = 1090000 // Maybe 1090000?
      } else if (m == 102000003) { // Warrior's Sanctuary
         npcId = 1022000
      } else if (m == 103000003) { // Thieves' Hideout
         npcId = 1052001
      } else if (m == 100000201) { // Bowman Instructional School
         npcId = 1012100
      } else if (m == 101000003) { // Magic Library
         npcId = 1032001
      }

      if (npcId != 0) {
         pi.openNpc(npcId)
         return true
      }
   }
   return false
}