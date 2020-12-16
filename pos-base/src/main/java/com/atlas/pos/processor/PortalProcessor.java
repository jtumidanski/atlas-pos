package com.atlas.pos.processor;

import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.pos.BlockedPortalRegistry;
import com.atlas.pos.event.producer.ChangeMapCommandProducer;
import com.atlas.pos.event.producer.EnableActionsCommandProducer;
import com.atlas.pos.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import rest.DataContainer;

public final class PortalProcessor {
   private PortalProcessor() {
   }

   public static List<Portal> getMapPortals(int mapId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .getRestClient(PortalAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::dataList)
            .orElse(Collections.emptyList())
            .stream()
            .map(ModelFactory::createPortal)
            .collect(Collectors.toList());
   }

   public static Optional<Portal> getMapPortalByName(int mapId, String name) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .queryParam("name", name)
            .getRestClient(PortalAttributes.class)
            .getWithResponse()
            .result()
            .flatMap(DataContainer::data)
            .map(ModelFactory::createPortal);
   }

   public static Optional<Portal> getMapPortalById(int mapId, int portalId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .pathParam("portals", portalId)
            .getRestClient(PortalAttributes.class)
            .getWithResponse()
            .result()
            .flatMap(DataContainer::data)
            .map(ModelFactory::createPortal);
   }

   public static void enterPortal(int worldId, int channelId, int characterId, int mapId, int portalId) {
      getMapPortalById(mapId, portalId).ifPresent(portal -> enterPortal(worldId, channelId, characterId, mapId, portal));
   }

   protected static boolean isSpawnPoint(Portal portal) {
      return (portal.type() == 0 || portal.type() == 1) && portal.targetMap() == 999999999;
   }

   protected static void enterPortal(int worldId, int channelId, int characterId, int mapId, Portal portal) {
      if (BlockedPortalRegistry.getInstance().isBlocked(characterId, portal.scriptName())) {
         EnableActionsCommandProducer.send(worldId, channelId, characterId);
         return;
      }

      boolean changed = false;
      if (portal.scriptName() != null && !portal.scriptName().equals("")) {
         changed = PortalScriptProcessor.getInstance().executePortalScript(worldId, channelId, characterId, mapId, portal);
      } else if (portal.targetMap() != 999999999) {
         //if (!(chr.getChalkboard() != null && GameConstants.isFreeMarketRoom(getTargetMapId()))) {
         // fallback for missing portals - no real life case anymore - interesting for not implemented areas
         Portal toPortal = getMapPortalByName(portal.targetMap(), portal.target())
               .orElse(getMapPortalById(portal.targetMap(), 0).orElseThrow());

         ChangeMapCommandProducer.changeMap(worldId, channelId, characterId, portal.targetMap(), toPortal.id());

         changed = true;
         //} else {
         //   MessageBroadcaster.getInstance().sendServerNotice(chr, ServerNoticeType.PINK_TEXT, I18nMessage.from("CANNOT_ENTER_MAP_WITH_CHALKBOARD_OPENED"));
         //}
      }

      if (!changed) {
         EnableActionsCommandProducer.send(worldId, channelId, characterId);
      }
   }
}
