package com.atlas.pos.script.portal


import scripting.portal.PortalPlayerInteraction

boolean enter(PortalPlayerInteraction pi) {
   pi.playPortalSound(); pi.warp(670010300, 0)
   return true
}