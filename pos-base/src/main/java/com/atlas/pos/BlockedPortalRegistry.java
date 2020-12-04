package com.atlas.pos;

import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

public class BlockedPortalRegistry {
   private static final Object lock = new Object();

   private static volatile BlockedPortalRegistry instance;

   private final Map<Integer, Set<String>> blockedPortalsPerCharacter;

   public static BlockedPortalRegistry getInstance() {
      BlockedPortalRegistry result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new BlockedPortalRegistry();
               instance = result;
            }
         }
      }
      return result;
   }

   private BlockedPortalRegistry() {
      blockedPortalsPerCharacter = new ConcurrentHashMap<>();
   }

   public boolean isBlocked(int characterId, String scriptName) {
      if (!blockedPortalsPerCharacter.containsKey(characterId)) {
         return false;
      }
      return blockedPortalsPerCharacter.get(characterId).contains(scriptName);
   }

   public void addBlockedPortal(int characterId, String scriptName) {
      if (!blockedPortalsPerCharacter.containsKey(characterId)) {
         blockedPortalsPerCharacter.put(characterId, new HashSet<>());
      }
      blockedPortalsPerCharacter.get(characterId).add(scriptName);
   }
}
