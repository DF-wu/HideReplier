package dfder.hidereplyer.test;

import com.google.gson.Gson;
import com.google.gson.internal.$Gson$Preconditions;
import dfder.hidereplyer.Entity.DiscordMessage;
import dfder.hidereplyer.Entity.DiscordWebhook;
import dfder.hidereplyer.Entity.Embedobj;

import java.awt.*;
import java.io.IOException;
import java.lang.reflect.Field;
import java.sql.Time;
import java.time.Clock;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.util.EnumMap;

//
public class TestDiscordPost {
    public static void main(String[] args) throws IOException
    {
        postTest();
    }
    
    
    
    public static void postTest() throws IOException
    {
        DiscordMessage webhook = new DiscordMessage("https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA");
        webhook.setContent("i can not input \nchinese");
        webhook.setAvatarUrl("https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraight&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=BlazerShirt&eyeType=Default&eyebrowType=Default&mouthType=Default&skinColor=Light");
        webhook.setUsername("Custom Usernames!");
        Embedobj em = new Embedobj();
        em.makeEmbed("title",
                "WTFdes",
                "http://dfder.cf",
                null,
                new Embedobj.Footer("hello footer","http://dfder.cf"),
                new Embedobj.Thumbnail("https://kryptongta.com/images/kryptontitle2.png"),
                new Embedobj.Image("http://dfder.cf"),
                new Embedobj.Author("wtfAU", "https://kryptongta.com/images/kryptontitle2.png","https://kryptongta.com/images/kryptontitle2.png"),
                new Embedobj.Field("Field name","value", true));
        em.addField("motherfucker","操你媽逼", true);
        webhook.addEmbed(em);
        
        webhook.excute();
        
    }
    
    
    public static DiscordWebhook testCreateAWebhookData(int number){
        DiscordWebhook webhook = new DiscordWebhook("https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA");
        webhook.setContent("哈囉你好嗎");
        webhook.setAvatarUrl("https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraight&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=BlazerShirt&eyeType=Default&eyebrowType=Default&mouthType=Default&skinColor=Light");
        webhook.setUsername("Custom Usernames!");
        // webhook.setTts(true);
        webhook.addEmbed(new DiscordWebhook.EmbedObject()
                .setTitle("Title")
                .setDescription("This is a description")
                .setColor(Color.RED)
                .addField("1st Field", "Inline", true)
                .addField("2nd Field", "Inline", true)
                .addField("3rd Field", "No-Inline", false)
                .setThumbnail("https://kryptongta.com/images/kryptonlogo.png")
                .setFooter("Footer text", "https://kryptongta.com/images/kryptonlogodark.png")
                .setImage("https://kryptongta.com/images/kryptontitle2.png")
                .setAuthor("Author Name", "https://kryptongta.com", "https://kryptongta.com/images/kryptonlogowide.png")
                .setUrl("https://kryptongta.com"));
        webhook.addEmbed(new DiscordWebhook.EmbedObject()
                .setDescription(String.valueOf(number)));
        
        Gson gson = new Gson();
        String st = gson.toJson(webhook);
        System.out.println( st);
        
        return webhook;
    
    
        
        
    }
    
    public static void testoriginal() throws IOException
    {
        DiscordWebhook webhook = new DiscordWebhook(DiscordWebhook.defaultUrl);
        webhook.setContent("Any message!");
        webhook.setAvatarUrl("https://your.awesome/image.png");
        webhook.setUsername("Custom Usernames!");
        webhook.setTts(true);
        webhook.addEmbed(new DiscordWebhook.EmbedObject()
                .setTitle("Title")
                .setDescription("This is a description")
                .setColor(Color.RED)
                .addField("1st Field", "Inline", true)
                .addField("2nd Field", "Inline", true)
                .addField("3rd Field", "No-Inline", false)
                .setThumbnail("https://kryptongta.com/images/kryptonlogo.png")
                .setFooter("Footer text", "https://kryptongta.com/images/kryptonlogodark.png")
                .setImage("https://kryptongta.com/images/kryptontitle2.png")
                .setAuthor("Author Name", "https://kryptongta.com", "https://kryptongta.com/images/kryptonlogowide.png")
                .setUrl("https://kryptongta.com"));
        webhook.addEmbed(new DiscordWebhook.EmbedObject()
                .setDescription("Just another added embed object!"));
    
        Gson gson = new Gson();
        System.out.println(gson.toJson(webhook));
        webhook.execute(); //Handle exception
    }
   
}

