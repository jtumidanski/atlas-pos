package com.atlas.pos.event.producer;

import com.atlas.csrv.command.EnableActionsCommand;
import com.atlas.csrv.constant.EventConstants;
import com.atlas.kafka.KafkaProducerFactory;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;

public class EnableActionsCommandProducer {
   private static final Object lock = new Object();

   private static volatile EnableActionsCommandProducer instance;

   private final Producer<Long, EnableActionsCommand> producer;

   public static EnableActionsCommandProducer getInstance() {
      EnableActionsCommandProducer result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new EnableActionsCommandProducer();
               instance = result;
            }
         }
      }
      return result;
   }

   public EnableActionsCommandProducer() {
      producer = KafkaProducerFactory.createProducer("Portal Service", System.getenv("BOOTSTRAP_SERVERS"));
   }

   public void send(int worldId, int channelId, int characterId) {
      String topic = System.getenv(EventConstants.TOPIC_ENABLE_ACTIONS);
      long key = produceKey(worldId, channelId);
      producer.send(new ProducerRecord<>(topic, key, new EnableActionsCommand(characterId)));
   }

   protected Long produceKey(int worldId, int channelId) {
      return (long) ((worldId * 1000) + channelId);
   }
}
