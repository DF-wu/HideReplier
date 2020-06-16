package dfder.hidereplyer.controller;



import dfder.hidereplyer.Entity.DiscordMessage;
import dfder.hidereplyer.Entity.DiscordStoreData;
import dfder.hidereplyer.service.DiscordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping(value = "/HideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class Discordbotcontroller {
    
    @Autowired
    private DiscordService discordService;
    
    //post a post to discord
    @PostMapping("/discord")
    //@RequestMapping(produces = "application/json;charset=UTF-8")
    public ResponseEntity<DiscordMessage> postToDiscord(@RequestBody DiscordMessage discordMessage) throws IOException
    {
        
        return ResponseEntity.ok().body(discordService.postApost(discordMessage));
    }

    
    
    @GetMapping("/discord")
    public ResponseEntity<ArrayList<DiscordStoreData>> getDiscordMessage()
    {
        return ResponseEntity.ok().body(discordService.gethistorylist());
    }
    
}
