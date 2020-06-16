package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.DiscordMessage;
import dfder.hidereplyer.Entity.DiscordStoreData;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.Clock;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.ArrayList;

@Service
public class DiscordService {
    
    
    @Autowired
    private DiscordPostDataRepo repo;
    private String url = "https://discord.com/api/webhooks/REDACTED/REDACTED";
    
    // post a post to discord
    public DiscordMessage postApost(DiscordMessage discordMessage) throws IOException
    {
        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
    
        for (int i = 10; i > 0; i--)
        {
            DiscordStoreData discordStoreData =
                    new DiscordStoreData(TWtime,i, discordMessage);
            repo.insert(discordStoreData);
            discordMessage.excute();
        }
        
        return discordMessage;
    }
    
    public ArrayList<DiscordStoreData> gethistorylist(){
       return repo.findAll();
        
    }
    
    
    
    
    // -----------below for testing--------------
    
    
    public void testService() throws IOException
    {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }
    
    
    
    
    
 
}
