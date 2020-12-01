package com.atlas.pos;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import com.atlas.kafka.consumer.ConsumerBuilder;
import com.atlas.pos.command.PortalEnterCommand;
import com.atlas.pos.constant.EventConstants;
import com.atlas.pos.consumer.PortalEnterCommandConsumer;

public class Server {
   private static final ExecutorService executorService = Executors.newSingleThreadExecutor();

   public static void main(String[] args) {

      executorService.execute(
            new ConsumerBuilder<>("Portal Service", PortalEnterCommand.class)
                  .setBootstrapServers(System.getenv("BOOTSTRAP_SERVERS"))
                  .setTopic(System.getenv(EventConstants.TOPIC_ENTER_PORTAL))
                  .setHandler(new PortalEnterCommandConsumer())
                  .build()
      );
   }
}
