package com.atlas.pos.processor;

import com.app.rest.util.RestResponseUtil;
import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.pos.BlockedPortalRegistry;
import com.atlas.pos.event.producer.ChangeMapCommandProducer;
import com.atlas.pos.event.producer.EnableActionsCommandProducer;
import com.atlas.pos.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;
import rest.DataContainer;

import java.util.Collection;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;

public final class PortalProcessor {
   private PortalProcessor() {
   }

   public static CompletableFuture<List<Portal>> getMapPortals(int mapId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::dataList)
            .thenApply(Collection::stream)
            .thenApply(stream -> stream.map(ModelFactory::createPortal))
            .thenApply(stream -> stream.collect(Collectors.toList()));
   }

   public static CompletableFuture<Portal> getMapPortalByName(int mapId, String name) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .queryParam("name", name)
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::data)
            .thenApply(Optional::get)
            .thenApply(ModelFactory::createPortal);
   }

   public static CompletableFuture<Portal> getMapPortalById(int mapId, int portalId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .pathParam("portals", portalId)
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::data)
            .thenApply(Optional::get)
            .thenApply(ModelFactory::createPortal);
   }

   public static void enterPortal(int worldId, int channelId, int characterId, int mapId, int portalId) {
      getMapPortalById(mapId, portalId)
            .thenAccept(portal -> enterPortal(worldId, channelId, characterId, mapId, portal));
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
               .exceptionally(fn -> getMapPortalById(portal.targetMap(), 0).join())
               .join();

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
