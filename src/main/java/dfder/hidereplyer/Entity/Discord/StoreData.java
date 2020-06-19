package dfder.hidereplyer.Entity.Discord;


import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

@Document(collection = "DiscordPostCollection")
public class StoreData {
    @Id
    private String id;
    
    @Field
    private long timeStamp;
    @Field
    private int serialNumber;
    @Field
    private PostMessage discordMessage;
    @Field
    private String posterIp;
    @Field
    private String responseCode;
    
    public StoreData(String id, long timeStamp, int serialNumber, PostMessage discordMessage)
    {
        this.id = id;
        this.timeStamp = timeStamp;
        this.serialNumber = serialNumber;
        this.discordMessage = discordMessage;
    }
    
    public StoreData(long tWtimeSecont, int counter, String posterIp, PostMessage dm)
    {
        this.timeStamp = tWtimeSecont;
        this.serialNumber = counter;
        this.posterIp = posterIp;
        this.discordMessage = dm;
    }
    
    public StoreData()
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
    
    
    public PostMessage getDiscordMessage()
    {
        return discordMessage;
    }
    
    public void setDiscordMessage(PostMessage discordMessage)
    {
        this.discordMessage = discordMessage;
    }
    
    public String getPosterIp()
    {
        return posterIp;
    }
    
    public void setPosterIp(String posterIp)
    {
        this.posterIp = posterIp;
    }
    
    public String getResponseCode()
    {
        return responseCode;
    }
    
    public void setResponseCode(String responseCode)
    {
        this.responseCode = responseCode;
    }
}
