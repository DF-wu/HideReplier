package dfder.hidereplyer.Entity;


import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.time.LocalDateTime;

@Document(collection = "DiscordPostCollection")
public class DiscordStoreData {
    @Id
    private String id;
    @Field
    private LocalDateTime timeStamp;
    @Field
    private int serialNumber;
    @Field
    private DiscordMessage discordMessage;
    
    public DiscordStoreData(LocalDateTime timeStamp, int serialNumber, DiscordMessage discordMessage)
    {
        this.timeStamp = timeStamp;
        this.serialNumber = serialNumber;
        this.discordMessage = discordMessage;
    }
    
    
    public void SerialNumberPlus1()
    {
        serialNumber++;
    }
    
    
    public String getId()
    {
        return id;
    }
    
    public void setId(String id)
    {
        this.id = id;
    }
    
    public LocalDateTime getTimeStamp()
    {
        return timeStamp;
    }
    
    public void setTimeStamp(LocalDateTime timeStamp)
    {
        this.timeStamp = timeStamp;
    }
    
    public int getSerialNumber()
    {
        return serialNumber;
    }
    
    public void setSerialNumber(int serialNumber)
    {
        this.serialNumber = serialNumber;
    }
    
    public DiscordMessage getDiscordMessage()
    {
        return discordMessage;
    }
    
    public void setDiscordMessage(DiscordMessage discordMessage)
    {
        this.discordMessage = discordMessage;
    }
}
