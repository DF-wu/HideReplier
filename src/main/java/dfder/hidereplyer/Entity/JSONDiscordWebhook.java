package dfder.hidereplyer.Entity;

import org.springframework.data.mongodb.core.mapping.Document;

import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.ZonedDateTime;

@Document(collection = "DiscordCollection")
public class JSONDiscordWebhook {
    private String id;
    
    
    private String content;
    private String botName;
    private LocalDateTime time;
    private String ip;
    
  
    public DiscordWebhook toDW(){
        DiscordWebhook dw = new DiscordWebhook(DiscordWebhook.defaultUrl);
        dw.setContent(content);
        dw.setUsername(botName);
        dw.setTimeStamp(time);
        return dw;
    }
    
    public String getId()
    {
        return id;
    }
    
    public void setId(String id)
    {
        this.id = id;
    }
    
    public String getContent()
    {
        return content;
    }
    
    public void setContent(String content)
    {
        this.content = content;
    }
    
    public String getBotName()
    {
        return botName;
    }
    
    public void setBotName(String botName)
    {
        this.botName = botName;
    }
    
    public LocalDateTime getTime()
    {
        return time;
    }
    
    public void setTime(LocalDateTime time)
    {
        this.time = time;
    }
}
