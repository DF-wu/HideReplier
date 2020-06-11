package dfder.hidereplyer.controller;


import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.service.DiscordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;

@RestController
@RequestMapping(value = "/HideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class Discordbotcontroller {
    
    @Autowired
    private DiscordService discordService;
    private DiscordPostDataRepo repo;
    
    //post a post to discord
    @PostMapping("/discord")
    public void postToDiscord(){
    
    }

    @GetMapping("/discord")
    public ResponseEntity<ArrayList<DiscordWebhook>> test()
    {
        return ResponseEntity.ok().body(discordService.testRepo());
    }
    
}
