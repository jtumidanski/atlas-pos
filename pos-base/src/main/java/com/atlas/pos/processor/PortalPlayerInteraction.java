package com.atlas.pos.processor;

import com.atlas.pos.model.Portal;

public class PortalPlayerInteraction {
   private final int characterId;

   private final Portal portal;

   public PortalPlayerInteraction(int characterId, Portal portal) {
      this.characterId = characterId;
      this.portal = portal;
   }
}
