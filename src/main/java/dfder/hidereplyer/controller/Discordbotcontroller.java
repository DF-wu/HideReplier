package dfder.hidereplyer.controller;



import dfder.hidereplyer.Entity.Discord.PostMessage;
import dfder.hidereplyer.Entity.Discord.StoreData;
import dfder.hidereplyer.Entity.RecivedJSONofPostMessage;
import dfder.hidereplyer.service.DiscordService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.ArrayList;


@RestController
@RequestMapping(value = "/HideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class Discordbotcontroller {
    
    @Autowired
    private DiscordService discordService;
    
    //post a post to discord
    @PostMapping(value = "/discord", produces = MediaType.APPLICATION_JSON_VALUE)
    //@RequestMapping(produces = "application/json;charset=UTF-8")
    public ResponseEntity<PostMessage> postToDiscord(@RequestBody RecivedJSONofPostMessage recivedJSONofPostMessage) throws IOException
    {
        //System.out.println(recivedJSONofPostMessage);
        //return ResponseEntity.ok().body(discordService.initTheWorld(recivedJSONofPostMessage));
        return ResponseEntity.ok().body(discordService.postApost(recivedJSONofPostMessage));
    }

    @GetMapping("/apiKey")
    public ResponseEntity<String> getApikey(){
        return ResponseEntity.ok().body(discordService.getTenorApiKey());
    }
    
    
    @GetMapping("/discord")
    public ResponseEntity<ArrayList<StoreData>> getDiscordMessage()
    {
        return ResponseEntity.ok().body(discordService.gethistorylist());
    }
    
    @GetMapping("/discord/version")
    public ResponseEntity<String> getVersion(){
        return ResponseEntity.ok().body(discordService.getVersion());
    }
    
}
