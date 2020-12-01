package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   EventManager elevator = pi.getEventManager("Elevator")
   if (elevator == null) {
      pi.sendPinkNotice("ELEVATOR_MAINTENANCE")
   } else if (elevator.getProperty(pi.getMapId() == 222020100 ? ("goingUp") : ("goingDown")) == "false") {
      pi.playPortalSound(); pi.warp(pi.getMapId() == 222020100 ? 222020110 : 222020210, 0)
      return true
   } else if (elevator.getProperty(pi.getMapId() == 222020100 ? ("goingUp") : ("goingDown")) == "true") {
      pi.sendPinkNotice("ELEVATOR_MOVING")
   }
   return false
}