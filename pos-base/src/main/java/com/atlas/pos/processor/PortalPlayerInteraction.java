package com.atlas.pos.processor;

import java.util.Collections;
import java.util.Optional;

import com.atlas.cos.attribute.CharacterAttributes;
import com.atlas.csrv.rest.attribute.InstructionAttributes;
import com.atlas.csrv.rest.builder.InstructionAttributesBuilder;
import com.atlas.pos.BlockedPortalRegistry;
import com.atlas.pos.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import builder.ResultObjectBuilder;
import rest.DataBody;
import rest.DataContainer;

public class PortalPlayerInteraction {
   private final int worldId;

   private final int channelId;

   private final int characterId;

   private final int mapId;

   private final Portal portal;

   public PortalPlayerInteraction(int worldId, int channelId, int characterId, int mapId, Portal portal) {
      this.worldId = worldId;
      this.channelId = channelId;
      this.characterId = characterId;
      this.mapId = mapId;
      this.portal = portal;
   }

   public int getCharacterId() {
      return characterId;
   }

   public int getMapId() {
      return mapId;
   }

   public Portal getPortal() {
      return portal;
   }

   public void playPortalSound() {
   }

   public void warp(int mapId) {
   }

   public void warp(int mapId, int portalId) {
   }

   public void warp(int mapId, String portalName) {
   }

   public void showInstruction(String message, int width, int height) {
      UriBuilder.service(RestService.CHANNEL)
            .pathParam("worlds", worldId)
            .pathParam("channels", channelId)
            .pathParam("characters", characterId)
            .path("instructions")
            .getRestClient()
            .create(new ResultObjectBuilder(InstructionAttributes.class, 0)
                  .setAttribute(
                        new InstructionAttributesBuilder()
                              .setMessage(message)
                              .setWidth(width)
                              .setHeight(height)
                  )
                  .inputObject()
            );
   }

   /**
    * Returns the mapId for the saved location.
    *
    * @param id
    * @return
    */
   public int getSavedLocation(String id) {
      return 0;
   }

   /**
    * Sends a pink notice to the character
    *
    * @param token
    */
   public void sendPinkNotice(String token) {
   }

   public void sendPinkNotice(String token, Object... args) {
   }

   public void blockPortal() {
      if (portal.scriptName() != null && !BlockedPortalRegistry.getInstance().isBlocked(characterId, portal.scriptName())) {
         BlockedPortalRegistry.getInstance().addBlockedPortal(characterId, portal.scriptName());
         EnableActionsProcessor.getInstance().send(worldId, channelId, characterId);
      }
   }

   public boolean containsAreaInfo(Short aShort, String s) {
      return false;
   }

   public void updateAreaInfo(Short aShort, String s) {
   }

   public void showInfo(String s) {
   }

   public void showIntro(String s) {
   }

   public void playSound(String s) {
   }

   public boolean isQuestStarted(int questId) {
      return false;
   }

   public void teachSkill(Integer integer1, Byte aByte1, Byte aByte2, Integer integer2) {
   }

   public boolean isQuestCompleted(int questId) {
      return false;
   }

   public boolean hasItem(int itemId) {
      return hasItem(itemId, 1);
   }

   public boolean hasItem(int itemId, int quantity) {
      return false;
   }

   public void isMorphed(int mobId) {
      //ch.getBuffSource(MapleBuffStat.MORPH) == 2210005
   }

   public void gainItem(int itemId, short quantity) {
   }

   public int canHold(int itemId) {
      return canHold(itemId, 1);
   }

   public int canHold(int itemId, int quantity) {
      return 0;
   }

   public void earnTitle(String s) {
   }

   public void lockUI() {
   }

   public void openNpc(Integer integer, String s) {
   }

   public void openNpc(Integer integer) {

   }

   public void setDirectionStatus(Boolean directionStatus) {
   }

   public void spawnNPC(Integer integer1, Integer integer2, Integer integer3, Integer integer4, Integer integer5,
                        Boolean aBoolean) {
   }

   public void setNPCValue(Integer integer, String s) {
   }

   public void updateInfo(String s1, String s2) {
   }

   public void sendDirectionInfo(Integer integer1, Integer integer2) {
   }

   public void unlockUI() {
   }

   public void enableActions() {
   }

   public void forceCompleteQuest(Integer integer) {
   }

   public void changeMusic(String s) {
   }

   public boolean isQuestActive(int questId) {
      return false;
   }

   public Integer getJobId() {
      return 0;
   }

   public void removeAll(int itemId) {
   }

   public int getGender() {
      return 0;
   }

   public void forceStartQuest(int questId) {
   }

   public void giveCharacterExp(int amount) {
   }

   public void runMapScript() {
   }

   public int getMarketPortalId(int mapId) {
      return 0;
   }

   public void saveLocation(String id) {
   }

   public int getTeam() {
      return 0;
   }

   public void showEffect(String s) {
   }

   public void mapEffect(String s) {
   }

   public void useItem(int itemId) {
   }

   public Optional<Object> getParty() {
      return Optional.empty();
   }

   public boolean isEventLeader() {
      return false;
   }

   public void talkGuide(String s) {
   }

   public void removeGuide() {
   }

   public void showInfoText(String s) {
   }

   /**
    * Counts items for character.
    *
    * @param itemId
    * @return
    */
   public int countItem(int itemId) {
      return 0;
   }

   public void cancelItem(Integer integer) {
   }

   public boolean hasLevel30Character() {
      CharacterAttributes attributes = UriBuilder.service(RestService.CHARACTER)
            .pathParam("characters", characterId)
            .getRestClient(CharacterAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getData)
            .map(DataBody::getAttributes)
            .orElseThrow();
      return UriBuilder.service(RestService.CHARACTER)
            .path("characters")
            .queryParam("accountId", attributes.accountId())
            .queryParam("worldId", attributes.worldId())
            .getRestClient(CharacterAttributes.class)
            .getWithResponse()
            .result()
            .map(DataContainer::getDataAsList)
            .orElse(Collections.emptyList())
            .stream()
            .anyMatch(body -> body.getAttributes().level() >= 30);
   }

   public void spawnGuide() {
   }

   /**
    * Gets the players level.
    *
    * @return
    */
   public int getLevel() {
      return 1;
   }

   public void guideHint(Integer integer) {
   }
}
