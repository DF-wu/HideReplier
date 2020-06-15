package dfder.hidereplyer.test;

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

//
public class TestDiscordPost {
    public static void main(String[] args) throws IOException
    {
        DiscordMessage dm = new DiscordMessage("https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA");
        dm.setContent("1\n\n\n1");
        dm.setUsername("hello");
        Embedobj em = new Embedobj();
        em.addField("n","v",true);
        dm.addEmbeds(em);
        dm.excute();
    }
    
    
    
    public static void postTest() throws IOException
    {
        DiscordMessage webhook = new DiscordMessage("https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA");
        webhook.setContent("i can not input \nchinese");
        webhook.setAvatarUrl("https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraight&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=BlazerShirt&eyeType=Default&eyebrowType=Default&mouthType=Default&skinColor=Light");
        webhook.setUsername("Custom Usernames!");
        // webhook.setTts(true);
        Embedobj em = new Embedobj();
        Embedobj em2 = new Embedobj();

        em.makeEmbed("title",
                "wtf",
                "https://imgur.com/a/8fCDVVb",
                Color.cyan,
                new Embedobj.Footer("wtf","https://imgur.com/a/8fCDVVb"),
                new Embedobj.Thumbnail("wtfthumbnail"),
                new Embedobj.Image("https://imgur.com/a/8fCDVVb"),
                new Embedobj.Author("authorDF","https://imgur.com/a/8fCDVVb","https://imgur.com/a/8fCDVVb"),
                new Embedobj.Field("str","val",true));
    
        System.out.println(em2.getColor());
        
        webhook.addEmbed(new Embedobj()
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
        webhook.execute(); //Handle exception
        
    }
    
    
    public DiscordWebhook testCreateAWebhookData(int number){
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
        
        return webhook;
    }
    
    
   
}

