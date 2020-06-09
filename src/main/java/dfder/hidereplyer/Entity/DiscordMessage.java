package dfder.hidereplyer.Entity;

import java.util.ArrayList;
import java.util.List;

public class DiscordMessage {
    private final String url;
    private String content;
    private String username;
    private String avatarUrl;
    private boolean tts;
    private List<DiscordWebhook.EmbedObject> embeds = new ArrayList<>();
    
    public DiscordMessage(String url)
    {
        this.url = url;
    }
    
    public String getContent()
    {
        return content;
    }
    
    public void setContent(String content)
    {
        this.content = content;
    }
    
    public String getUsername()
    {
        return username;
    }
    
    public void setUsername(String username)
    {
        this.username = username;
    }
    
    public String getAvatarUrl()
    {
        return avatarUrl;
    }
    
    public void setAvatarUrl(String avatarUrl)
    {
        this.avatarUrl = avatarUrl;
    }
    
    public boolean isTts()
    {
        return tts;
    }
    
    public void setTts(boolean tts)
    {
        this.tts = tts;
    }
}
