package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.ArrayList;

@Service
public class DiscordService {
    
    private DiscordWebhook discordWebhook;
    
    @Autowired
    private DiscordPostDataRepo repo;

    
    public void postApost(){
        // post a post to discord
        
        
    }
    
    
    public void testService() throws IOException
    {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }
    
    public ArrayList<DiscordWebhook> testRepo(){
        ArrayList<DiscordWebhook> dwA = new ArrayList<>();
        for (int i = 0; i < 10000; i++)
        {
            TestDiscordPost tds = new TestDiscordPost();
            DiscordWebhook dw = tds.testCreateAWebhookData(i);
            dwA.add(dw);
            repo.insert(dw);
        }
        return dwA;
    }
    
    
    
 
}
