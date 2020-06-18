package dfder.hidereplyer.Entity;


import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.awt.*;
import java.time.LocalDateTime;

@Document(collection = "DiscordPostCollection")
public class DiscordStoreData {
    @Id
    private String id;
    

    @Field
    private long timeStamp;
    @Field
    private int serialNumber;
    @Field
    private DiscordMessage discordMessage;
    

    
    public DiscordStoreData(String id, long timeStamp, int serialNumber, DiscordMessage discordMessage)
    {
        this.id = id;
        this.timeStamp = timeStamp;
        this.serialNumber = serialNumber;
        this.discordMessage = discordMessage;
    }
    
    public DiscordStoreData(long tWtimeSecont, int counter, DiscordMessage dm)
    {
        this.timeStamp = tWtimeSecont;
        this.serialNumber = counter;
        this.discordMessage = dm;
    }
    
    public DiscordStoreData()
    {
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
    
    public int getSerialNumber()
    {
        return serialNumber;
    }
    
    public void setSerialNumber(int serialNumber)
    {
        this.serialNumber = serialNumber;
    }
    
    public long getTimeStamp()
    {
        return timeStamp;
    }
    
    public void setTimeStamp(long timeStamp)
    {
        this.timeStamp = timeStamp;
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
