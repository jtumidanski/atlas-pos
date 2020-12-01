package com.atlas.pos.script.portal

import com.atlas.pos.processor.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
    pi.showInstruction("You can move by using the arrow keys.", 250, 5)
    return true
}