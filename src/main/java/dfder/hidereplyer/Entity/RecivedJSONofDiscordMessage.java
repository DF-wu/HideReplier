package dfder.hidereplyer.Entity;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import org.springframework.data.annotation.Id;
import org.springframework.data.repository.NoRepositoryBean;

import java.util.HashMap;
import java.util.Map;

public class RecivedJSONofDiscordMessage extends DiscordMessage{
    
    public void setExtras(Map<String, Object> extras)
    {
        this.extras = extras;
    }
    
    public Map<String,Object> extras = new HashMap<>();
    
    @JsonAnySetter
    public void addExtras(String key, Object value) {
        this.extras.put(key, value);
    }
    
    @JsonAnyGetter
    public Object getExtra(String key)
    {
        return this.extras.get(key);
    }
}
