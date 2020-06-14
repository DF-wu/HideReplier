package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.Entity.JSONDiscordWebhook;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.apache.tomcat.jni.Local;
import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.RequestBody;

import java.io.IOException;
import java.time.Clock;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.ArrayList;
import java.util.Map;

@Service
public class DiscordService {
    
    private DiscordWebhook discordWebhook;
    
    @Autowired
    private DiscordPostDataRepo repo;
    private String url = "https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA";
    
    // post a post to discord
    public JSONDiscordWebhook postApost(Map<String, Object> jsonObject) throws IOException
    {
        
        JSONDiscordWebhook json = new JSONDiscordWebhook();
        json.setBotName((String) jsonObject.get("username"));
        
        json.setContent((String) jsonObject.get("content"));
        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
        json.setTime(TWtime);
        
        repo.insert(json);
        DiscordWebhook dw = json.toDW();
        dw.execute();
        return json;
        
    }
    
    
    
    
    
    // -----------below for testing--------------
    
    
    public void testService() throws IOException
    {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }
    
    
    
    
    
 
}
