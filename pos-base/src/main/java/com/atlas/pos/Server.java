package com.atlas.pos;

import com.atlas.kafka.consumer.SimpleEventConsumerFactory;
import com.atlas.pos.event.consumer.PortalEnterCommandConsumer;

public class Server {
   public static void main(String[] args) {
      SimpleEventConsumerFactory.create(new PortalEnterCommandConsumer());
   }
}
