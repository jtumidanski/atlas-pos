package com.atlas.pos.processor;

import com.atlas.cos.attribute.CharacterAttributes;
import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.pos.model.Character;
import com.atlas.pos.model.Portal;

import rest.DataBody;

public final class ModelFactory {
   private ModelFactory() {
   }

   public static Portal createPortal(DataBody<PortalAttributes> body) {
      return new Portal(Integer.parseInt(body.getId()),
            body.getAttributes().name(),
            body.getAttributes().target(),
            body.getAttributes().type(),
            body.getAttributes().x(),
            body.getAttributes().y(),
            body.getAttributes().targetMap(),
            body.getAttributes().scriptName()
      );
   }

   public static Character createCharacter(DataBody<CharacterAttributes> body) {
      return new Character(Integer.parseInt(body.getId()),
            body.getAttributes().level(),
            body.getAttributes().x(),
            body.getAttributes().y());
   }
}
