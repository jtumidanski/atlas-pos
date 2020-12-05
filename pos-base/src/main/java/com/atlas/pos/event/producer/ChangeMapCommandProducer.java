package com.atlas.pos.event.producer;

import com.atlas.cos.command.ChangeMapCommand;
import com.atlas.cos.constant.EventConstants;
import com.atlas.kafka.KafkaProducerFactory;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;

public class ChangeMapCommandProducer {
   private static final Object lock = new Object();

   private static volatile ChangeMapCommandProducer instance;

   private final Producer<Long, ChangeMapCommand> producer;

   public static ChangeMapCommandProducer getInstance() {
      ChangeMapCommandProducer result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new ChangeMapCommandProducer();
               instance = result;
            }
         }
      }
      return result;
   }

   private ChangeMapCommandProducer() {
      producer = KafkaProducerFactory.createProducer("Portal Service", System.getenv("BOOTSTRAP_SERVERS"));
   }

   public void changeMap(int worldId, int channelId, int characterId, int mapId, int portalId) {
      String topic = System.getenv(EventConstants.TOPIC_CHANGE_MAP_COMMAND);
      long key = produceKey(worldId, channelId);
      producer.send(new ProducerRecord<>(topic, key, new ChangeMapCommand(worldId, channelId, characterId, mapId, portalId)));
   }

   protected Long produceKey(int worldId, int channelId) {
      return (long) ((worldId * 1000) + channelId);
   }
}
