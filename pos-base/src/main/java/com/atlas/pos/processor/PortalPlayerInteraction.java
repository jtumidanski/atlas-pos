package com.atlas.pos.processor;

import builder.ResultObjectBuilder;
import com.app.rest.util.RestResponseUtil;
import com.atlas.cos.attribute.CharacterAttributes;
import com.atlas.csrv.rest.attribute.InstructionAttributes;
import com.atlas.csrv.rest.builder.InstructionAttributesBuilder;
import com.atlas.pos.BlockedPortalRegistry;
import com.atlas.pos.event.producer.ChangeMapCommandProducer;
import com.atlas.pos.event.producer.EnableActionsCommandProducer;
import com.atlas.pos.model.Character;
import com.atlas.pos.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;
import rest.DataBody;
import rest.DataContainer;

import java.awt.*;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;

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

   /**
    * Gets the character id interacting with the portal.
    *
    * @return the character identifier
    */
   public int getCharacterId() {
      return characterId;
   }

   /**
    * Gets the map id the portal resides in.
    *
    * @return the map identifier
    */
   public int getMapId() {
      return mapId;
   }

   /**
    * Gets the portal being interacted with.
    *
    * @return the portal
    */
   public Portal getPortal() {
      return portal;
   }

   public void playPortalSound() {
   }

   /**
    * Warps the character to the provided map, to the first portal.
    *
    * @param mapId the map identifier
    */
   public void warp(int mapId) {
      warp(mapId, 0);
   }

   /**
    * Warps the character to the provided map, to the portal specified.
    *
    * @param mapId    the map identifier
    * @param portalId the portal identifier
    */
   public void warp(int mapId, int portalId) {
      ChangeMapCommandProducer.changeMap(worldId, channelId, characterId, mapId, portalId);
   }

   /**
    * Warps the character to the provided map, to the portal specified.
    *
    * @param mapId      the map identifier
    * @param portalName the portal name
    */
   public void warp(int mapId, String portalName) {
      PortalProcessor.getMapPortalByName(mapId, portalName)
            .thenApply(Portal::id)
            .exceptionally(fn -> 0)
            .thenAccept(id -> warp(mapId, id));
   }

   /**
    * Shows an instruction to the character
    *
    * @param message the message
    * @param width   the width of the instruction
    * @param height  the height of the instruction
    */
   public void showInstruction(String message, int width, int height) {
      UriBuilder.service(RestService.CHANNEL)
            .pathParam("worlds", worldId)
            .pathParam("channels", channelId)
            .pathParam("characters", characterId)
            .path("instructions")
            .getAsyncRestClient(InstructionAttributes.class)
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
    * @param type the type of location to retrieve.
    * @return the map identifier
    */
   public int getSavedLocation(String type) {
      int location = CharacterProcessor.getSavedLocation(characterId, type);
      CharacterProcessor.saveLocation(characterId, type, 0, 0);
      return location;
   }

   public void sendPinkNotice(String token) {
   }

   public void sendPinkNotice(String token, Object... args) {
   }

   /**
    * Marks a portal as blocked to prevent handling again.
    */
   public void blockPortal() {
      if (portal.scriptName() != null && !BlockedPortalRegistry.getInstance().isBlocked(characterId, portal.scriptName())) {
         BlockedPortalRegistry.getInstance().addBlockedPortal(characterId, portal.scriptName());
         EnableActionsCommandProducer.send(worldId, channelId, characterId);
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

   /**
    * Gets the job of the character.
    *
    * @return the job id
    */
   public Integer getJobId() {
      return CharacterProcessor.getCharacter(characterId)
            .thenApply(Character::jobId)
            .join();
   }

   public void removeAll(int itemId) {
   }

   /**
    * Gets the gender of the character.
    *
    * @return 0 if male, 1 if female
    */
   public int getGender() {
      return CharacterProcessor.getCharacter(characterId)
            .thenApply(Character::gender)
            .join();
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

   /**
    * Saves a location of interest for the character.
    *
    * @param type the type of location
    */
   public void saveLocation(String type) {
      CompletableFuture<Point> fromFuture = CharacterProcessor.getCharacter(characterId)
            .thenApply(character -> new Point(character.x(), character.y()))
            .exceptionally(fn -> new Point(0, 0));
      CompletableFuture<List<Portal>> portalsFuture = PortalProcessor.getMapPortals(mapId);

      fromFuture.thenCombine(portalsFuture, this::findClosest)
            .thenApply(Optional::get)
            .thenApply(Portal::id)
            .exceptionally(fn -> 0)
            .thenAccept(id -> CharacterProcessor.saveLocation(characterId, type, mapId, id));
   }

   /**
    * Finds the portal closest to the reference point.
    *
    * @param from    the reference point
    * @param portals the portals to consider
    * @return the closest portal (if one exists).
    */
   protected Optional<Portal> findClosest(Point from, List<Portal> portals) {
      return portals.stream()
            .filter(PortalProcessor::isSpawnPoint)
            .min((o1, o2) -> compareDistanceFromPoint(from, o1, o2));
   }

   /**
    * Compares the distance between the two points, from a point of interest.
    *
    * @param from the point of interest
    * @param o1   the first point
    * @param o2   the second point
    * @return the value 0 if o1 is equal distance to the point of interest as d2; a value less than 0 if o1 is closer
    * than o2 to the point of interest; and a value greater than 0 if o1 is further than d2 from the point of interest.
    */
   protected static int compareDistanceFromPoint(Point from, Portal o1, Portal o2) {
      double o1Distance = new Point(o1.x(), o1.y()).distanceSq(from);
      double o2Distance = new Point(o2.x(), o2.y()).distanceSq(from);
      return Double.compare(o1Distance, o2Distance);
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

   public int countItem(int itemId) {
      return 0;
   }

   public void cancelItem(Integer integer) {
   }

   /**
    * Determines if the account interacting with the portal, has a level 30 character.
    *
    * @return true if a level 30 character or greater exists for the account
    */
   public boolean hasLevel30Character() {
      return UriBuilder.service(RestService.CHARACTER)
            .pathParam("characters", characterId)
            .getAsyncRestClient(CharacterAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::data)
            .thenApply(Optional::get)
            .thenApply(DataBody::getAttributes)
            .thenCompose(attributes -> getAccountCharacters(attributes.accountId(), attributes.worldId()))
            .thenApply(this::hasLevelThirtyCharacter)
            .join();
   }

   protected boolean hasLevelThirtyCharacter(List<DataBody<CharacterAttributes>> characters) {
      return characters.stream()
            .map(DataBody::getAttributes)
            .map(CharacterAttributes::level)
            .anyMatch(level -> level >= 30);
   }

   protected CompletableFuture<List<DataBody<CharacterAttributes>>> getAccountCharacters(int accountId, int worldId) {
      return UriBuilder.service(RestService.CHARACTER)
            .path("characters")
            .queryParam("accountId", accountId)
            .queryParam("worldId", worldId)
            .getAsyncRestClient(CharacterAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::dataList);
   }

   public void spawnGuide() {
   }

   /**
    * Gets the level of the character interacting with the portal.
    *
    * @return the level of the character
    */
   public int getLevel() {
      return CharacterProcessor.getCharacter(characterId)
            .thenApply(Character::level)
            .exceptionally(fn -> 1)
            .join();
   }

   public void guideHint(Integer integer) {
   }
}
