package com.atlas.pos.processor;

import com.atlas.cos.command.ChangeMapCommand;
import com.atlas.cos.constant.EventConstants;
import com.atlas.kafka.KafkaProducerFactory;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;

public class ChangeMapProcessor {
   private static final Object lock = new Object();

   private static volatile ChangeMapProcessor instance;

   private final Producer<Long, ChangeMapCommand> producer;

   public static ChangeMapProcessor getInstance() {
      ChangeMapProcessor result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new ChangeMapProcessor();
               instance = result;
            }
         }
      }
      return result;
   }

   private ChangeMapProcessor() {
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
