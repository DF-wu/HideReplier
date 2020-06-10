package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.repo.DiscordMessageRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class DiscordService {
    
    private DiscordWebhook discordWebhook;
    
    @Autowired
    private DiscordMessageRepo repo;
    
    
    public void postApost(){
        // post a post to discord
        
        
    }
    
    
    public void testService()
    {
    
    }
    
    
}
