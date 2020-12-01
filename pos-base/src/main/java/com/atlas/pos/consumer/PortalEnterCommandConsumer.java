package com.atlas.pos.consumer;

import com.atlas.kafka.consumer.ConsumerRecordHandler;
import com.atlas.pos.command.PortalEnterCommand;
import com.atlas.pos.processor.PortalProcessor;

public class PortalEnterCommandConsumer implements ConsumerRecordHandler<Long, PortalEnterCommand> {
   @Override
   public void handle(Long key, PortalEnterCommand command) {
      PortalProcessor.getInstance().enterPortal(command.worldId(), command.channelId(), command.characterId(), command.mapId(),
            command.portalId());
   }
}
