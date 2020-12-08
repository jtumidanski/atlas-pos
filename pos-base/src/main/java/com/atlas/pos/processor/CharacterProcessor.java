package com.atlas.pos.processor;

import java.util.Collections;
import java.util.Optional;

import com.atlas.cos.attribute.LocationAttributes;
import com.atlas.cos.attribute.LocationType;
import com.atlas.cos.builder.LocationAttributesBuilder;
import com.atlas.pos.model.Character;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import builder.ResultObjectBuilder;
import rest.DataBody;
import rest.DataContainer;

public final class CharacterProcessor {
   private CharacterProcessor() {
   }

   public static Optional<Character> getCharacter(int characterId) {
      return UriBuilder.service(RestService.CHARACTER)
            .pathParam("characters", characterId)
            .getRestClient(com.atlas.cos.attribute.CharacterAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getData)
            .map(ModelFactory::createCharacter);
   }

   public static void saveLocation(int characterId, String type, int mapId, int portalId) {
      UriBuilder.service(RestService.CHARACTER)
            .pathParam("characters", characterId)
            .path("locations")
            .getRestClient()
            .create(new ResultObjectBuilder(LocationAttributes.class, 0)
                  .setAttribute(
                        new LocationAttributesBuilder()
                              .setType(LocationType.valueOf(type))
                              .setMapId(mapId)
                              .setPortalId(portalId)
                  )
                  .inputObject()
            );
   }

   public static int getSavedLocation(int characterId, String type) {
      return UriBuilder.service(RestService.CHARACTER)
            .pathParam("characters", characterId)
            .path("locations")
            .queryParam("type", type)
            .getRestClient(LocationAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getDataAsList)
            .orElse(Collections.emptyList())
            .stream()
            .findFirst()
            .map(DataBody::getAttributes)
            .map(LocationAttributes::mapId)
            .orElse(0);
   }
}
