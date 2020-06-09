package dfder.hidereplyer.controller;


import dfder.hidereplyer.service.DiscordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/hideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class controller {
    
    @Autowired
    private DiscordService discordService;
    
    
    //post a post to discord
    @PostMapping("/discord")
    public void postToDiscord(){
    
    }


}
