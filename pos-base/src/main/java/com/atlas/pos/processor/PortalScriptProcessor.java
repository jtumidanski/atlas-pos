package com.atlas.pos.processor;

import com.atlas.pos.model.Portal;

import javax.script.*;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class PortalScriptProcessor {
   private static final Object lock = new Object();

   private static volatile PortalScriptProcessor instance;

   private final Map<String, ScriptEngine> scripts = new HashMap<>();

   private final ScriptEngineFactory sef;

   public static PortalScriptProcessor getInstance() {
      PortalScriptProcessor result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new PortalScriptProcessor();
               instance = result;
            }
         }
      }
      return result;
   }

   private PortalScriptProcessor() {
      sef = new ScriptEngineManager().getEngineByName("groovy").getFactory();
   }

   protected ScriptEngine getScriptEngine(String path) {
      ScriptEngine engine = sef.getScriptEngine();
      engine = evalPrerequisites(engine, getPrerequisites());
      return eval(engine, path);
   }

   protected ScriptEngine eval(ScriptEngine engine, String path) {
      File scriptFile = null;
      if (new File(path).exists()) {
         scriptFile = new File(path);
      }
      if (scriptFile == null) {
         return null;
      }

      try (FileReader fr = new FileReader(scriptFile)) {
         engine.eval(fr);
      } catch (final ScriptException | IOException t) {
         return null;
      }
      return engine;
   }

   protected String[] getPrerequisites() {
      return new String[0];
   }

   protected ScriptEngine evalPrerequisites(ScriptEngine engine, String... paths) {
      ScriptEngine primedEngine = engine;
      for (String path : paths) {
         primedEngine = eval(primedEngine, path);
         if (primedEngine == null) {
            return null;
         }
      }
      return primedEngine;
   }

   protected ScriptEngine getPortalScript(String scriptName) {
      String scriptPath = "/service/script/" + scriptName + ".groovy";
      ScriptEngine iv = scripts.get(scriptPath);
      if (iv != null) {
         return iv;
      }

      iv = getScriptEngine(scriptPath);
      if (iv == null) {
         return null;
      }
      scripts.put(scriptPath, iv);
      return iv;
   }

   public boolean executePortalScript(int worldId, int channelId, int characterId, int mapId, Portal portal) {
      try {
         ScriptEngine iv = getPortalScript(portal.scriptName());
         if (iv != null) {
            return (boolean) ((Invocable) iv).invokeFunction("enter", new PortalPlayerInteraction(worldId, channelId, characterId,
                  mapId,
                  portal));
         }
      } catch (Exception exception) {
         exception.printStackTrace();
      }
      return false;
   }
}
