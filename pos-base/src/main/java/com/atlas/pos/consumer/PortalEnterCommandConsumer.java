package com.atlas.pos.consumer;

import com.atlas.kafka.KafkaConsumerFactory;
import com.atlas.pos.command.PortalEnterCommand;
import com.atlas.pos.constant.EventConstants;
import com.atlas.pos.processor.PortalProcessor;
import org.apache.kafka.clients.consumer.Consumer;
import org.apache.kafka.clients.consumer.ConsumerRecords;

import java.time.Duration;
import java.util.Collections;

public class PortalEnterCommandConsumer implements Runnable {
   private final String bootstrapServers;

   private final String enterPortalTopic;

   public PortalEnterCommandConsumer(String bootstrapServers) {
      this.bootstrapServers = bootstrapServers;
      this.enterPortalTopic = System.getenv(EventConstants.TOPIC_ENTER_PORTAL);
   }

   @Override
   public void run() {
      Consumer<Long, PortalEnterCommand> consumer = KafkaConsumerFactory.createConsumer("World Registry",
            bootstrapServers, PortalEnterCommand.class);
      consumer.subscribe(Collections.singleton(enterPortalTopic));

      try (consumer) {
         while (true) {
            final ConsumerRecords<Long, PortalEnterCommand> consumerRecords = consumer.poll(Duration.ofMillis(1000));

            if (consumerRecords.count() > 0) {
               consumerRecords.forEach(record -> {
                  PortalEnterCommand command = record.value();
                  PortalProcessor.getInstance().enterPortal(command.characterId(), command.mapId(), command.portalId());
               });
               consumer.commitAsync();
            }
         }
      }
   }
}
