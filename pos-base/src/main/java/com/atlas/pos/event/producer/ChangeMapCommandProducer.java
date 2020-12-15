package com.atlas.pos.event.producer;

import com.atlas.cos.command.ChangeMapCommand;
import com.atlas.cos.constant.EventConstants;
import com.atlas.pos.EventProducerRegistry;

public final class ChangeMapCommandProducer {
   private ChangeMapCommandProducer() {
   }

   public static void changeMap(int worldId, int channelId, int characterId, int mapId, int portalId) {
      EventProducerRegistry.getInstance().send(ChangeMapCommand.class,
            EventConstants.TOPIC_CHANGE_MAP_COMMAND, worldId, channelId,
            new ChangeMapCommand(worldId, channelId, characterId, mapId, portalId));
   }
}
