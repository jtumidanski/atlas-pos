package com.atlas.pos.event.producer;

import com.atlas.csrv.command.EnableActionsCommand;
import com.atlas.csrv.constant.EventConstants;
import com.atlas.pos.EventProducerRegistry;

public final class EnableActionsCommandProducer {
   private EnableActionsCommandProducer() {
   }

   public static void send(int characterId) {
      EventProducerRegistry.getInstance().send(EventConstants.TOPIC_ENABLE_ACTIONS, characterId,
            new EnableActionsCommand(characterId));
   }
}
