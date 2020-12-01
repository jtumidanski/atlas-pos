package com.atlas.pos;

import com.atlas.pos.consumer.PortalEnterCommandConsumer;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class Server {
   private static final ExecutorService executorService = Executors.newSingleThreadExecutor();

   public static void main(String[] args) {
      String bootstrapServers = System.getenv("BOOTSTRAP_SERVERS");

      executorService.execute(new PortalEnterCommandConsumer(bootstrapServers));
   }
}
