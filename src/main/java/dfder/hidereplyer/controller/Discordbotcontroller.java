package dfder.hidereplyer.controller;


import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.Entity.JSONDiscordWebhook;
import dfder.hidereplyer.service.DiscordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Map;

@RestController
@RequestMapping(value = "/HideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class Discordbotcontroller {
    
    @Autowired
    private DiscordService discordService;
    
    //post a post to discord
    @PostMapping("/discord")
    //@RequestMapping(produces = "application/json;charset=UTF-8")
    public JSONDiscordWebhook postToDiscord(@RequestBody Map<String,Object> jsonObject) throws IOException
    {
        return discordService.postApost(jsonObject);
    }

    
    
}
