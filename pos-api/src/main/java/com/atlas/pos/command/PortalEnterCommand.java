package com.atlas.pos.command;

public record PortalEnterCommand(int worldId, int channelId, int mapId, int portalId, int characterId) {
}
