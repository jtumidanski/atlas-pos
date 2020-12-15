package com.atlas.pos.event.consumer;

import com.atlas.kafka.consumer.SimpleEventHandler;
import com.atlas.pos.command.PortalEnterCommand;
import com.atlas.pos.constant.EventConstants;
import com.atlas.pos.processor.PortalProcessor;
import com.atlas.pos.processor.TopicDiscoveryProcessor;

public class PortalEnterCommandConsumer implements SimpleEventHandler<PortalEnterCommand> {
   @Override
   public void handle(Long key, PortalEnterCommand command) {
      PortalProcessor.enterPortal(command.worldId(), command.channelId(), command.characterId(), command.mapId(),
            command.portalId());
   }

   @Override
   public Class<PortalEnterCommand> getEventClass() {
      return PortalEnterCommand.class;
   }

   @Override
   public String getConsumerId() {
      return "Portal Service";
   }

   @Override
   public String getBootstrapServers() {
      return System.getenv("BOOTSTRAP_SERVERS");
   }

   @Override
   public String getTopic() {
      return TopicDiscoveryProcessor.getTopic(EventConstants.TOPIC_ENTER_PORTAL);
   }
}
