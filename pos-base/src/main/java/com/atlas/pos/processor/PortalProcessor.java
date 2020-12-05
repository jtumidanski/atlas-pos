package com.atlas.pos.processor;

import java.util.Optional;

import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.pos.BlockedPortalRegistry;
import com.atlas.pos.event.producer.ChangeMapCommandProducer;
import com.atlas.pos.event.producer.EnableActionsCommandProducer;
import com.atlas.pos.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import rest.DataContainer;

public class PortalProcessor {
   private static final Object lock = new Object();

   private static volatile PortalProcessor instance;

   public static PortalProcessor getInstance() {
      PortalProcessor result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new PortalProcessor();
               instance = result;
            }
         }
      }
      return result;
   }

   public Optional<Portal> getMapPortalByName(int mapId, String name) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .queryParam("name", name)
            .getRestClient(PortalAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getData)
            .map(ModelFactory::createPortal);
   }

   public Optional<Portal> getMapPortalById(int mapId, int portalId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .pathParam("portals", portalId)
            .getRestClient(PortalAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getData)
            .map(ModelFactory::createPortal);
   }

   public void enterPortal(int worldId, int channelId, int characterId, int mapId, int portalId) {
      getMapPortalById(mapId, portalId).ifPresent(portal -> enterPortal(worldId, channelId, characterId, mapId, portal));
   }

   protected void enterPortal(int worldId, int channelId, int characterId, int mapId, Portal portal) {
      if (BlockedPortalRegistry.getInstance().isBlocked(characterId, portal.scriptName())) {
         EnableActionsCommandProducer.getInstance().send(worldId, channelId, characterId);
         return;
      }

      boolean changed = false;
      if (portal.scriptName() != null) {
         changed = PortalScriptProcessor.getInstance().executePortalScript(worldId, channelId, characterId, mapId, portal);
      } else if (portal.targetMap() != 999999999) {
         //if (!(chr.getChalkboard() != null && GameConstants.isFreeMarketRoom(getTargetMapId()))) {
         // fallback for missing portals - no real life case anymore - interesting for not implemented areas
         Portal toPortal = getMapPortalByName(portal.targetMap(), portal.target())
               .orElse(getMapPortalById(portal.targetMap(), 0).orElseThrow());

         ChangeMapCommandProducer.getInstance().changeMap(worldId, channelId, characterId, portal.targetMap(), toPortal.id());

         changed = true;
         //} else {
         //   MessageBroadcaster.getInstance().sendServerNotice(chr, ServerNoticeType.PINK_TEXT, I18nMessage.from("CANNOT_ENTER_MAP_WITH_CHALKBOARD_OPENED"));
         //}
      }

      if (!changed) {
         EnableActionsCommandProducer.getInstance().send(worldId, channelId, characterId);
      }
   }
}
