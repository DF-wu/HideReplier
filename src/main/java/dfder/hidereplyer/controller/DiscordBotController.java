package dfder.hidereplyer.controller;



import dfder.hidereplyer.Entity.Discord.PostMessage;
import dfder.hidereplyer.Entity.Discord.StoreData;
import dfder.hidereplyer.Entity.RecivedJSONofPostMessage;
import dfder.hidereplyer.service.DiscordService;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.ArrayList;


@RestController
@RequestMapping(value = "/HideBot" , produces = MediaType.APPLICATION_JSON_VALUE)
public class DiscordBotController {

    private final DiscordService discordService;

    public DiscordBotController(DiscordService discordService) {
        this.discordService = discordService;
    }

    @PostMapping(value = "/discord", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<PostMessage> postToDiscord(@RequestBody RecivedJSONofPostMessage postMessage) throws IOException {
        return ResponseEntity.ok().body(discordService.postAnonymousMessage(postMessage));
    }


    @GetMapping("/discord")
    public ResponseEntity<ArrayList<StoreData>> getDiscordMessage() {
        return ResponseEntity.ok().body(discordService.getHistoryList());
    }

    @GetMapping("/discord/version")
    public ResponseEntity<String> getVersion() {
        return ResponseEntity.ok().body(discordService.getVersion());
    }
}
