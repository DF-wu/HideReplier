package dfder.hidereplyer.test;

import com.google.gson.Gson;
import dfder.hidereplyer.Entity.DiscordMessage;

import dfder.hidereplyer.Entity.DiscordStoreData;
import dfder.hidereplyer.Entity.Embedobj;
import dfder.hidereplyer.resource.DiscordWebhook;

import java.awt.*;
import java.io.IOException;
import java.time.Clock;
import java.time.LocalDateTime;
import java.time.ZoneId;

//
public class TestDiscordPost {
    public static void main(String[] args) throws IOException
    {
        messagedPost();
    }
    
    
    
    public static void postTest() throws IOException
    {

    }
    
    
    public static DiscordWebhook testCreateAWebhookData(int number){
        DiscordWebhook webhook = new DiscordWebhook("http");
        webhook.setContent("哈囉你好嗎");
        webhook.setAvatarUrl("https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraight&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=BlazerShirt&eyeType=Default&eyebrowType=Default&mouthType=Default&skinColor=Light");
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
                .setDescription(String.valueOf(number)));
        
        Gson gson = new Gson();
        String st = gson.toJson(webhook);
        System.out.println( st);
        
        return webhook;
    
    
        
        
    }
    
    public static void testoriginal() throws IOException
    {
        DiscordWebhook webhook = new DiscordWebhook(DiscordMessage.defaultUrl);
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
    
    
    public static void messagedPost() throws IOException
    {
        //DiscordMessage webhook = new DiscordMessage("https://discordapp.com/api/webhooks/710112845567623238/_LxpGhvlK6Wp-LgavxRV1UlcsEdqnnznHXuA6d4v7YWUBCfHOBRPwCo2B7-ao0_3EbwV");
        DiscordMessage webhook = new DiscordMessage(DiscordMessage.defaultUrl);
        //webhook.setContent("ㄟ不是我沒有權限刪除阿....");
        webhook.setAvatarUrl("https://img.icons8.com/color/144/000000/drupal.png");
        webhook.setUsername("匿名機器人ㄎㄎ");
        //webhook.setTts(true);
        Embedobj em = new Embedobj();
        em.makeEmbed("機器人名字",//"不是我沒有權限刪除阿....",
                "靠北訊息 ",
                "https://img.icons8.com/color/288/000000/drupal.png",
                Color.YELLOW,
                null,//new Embedobj.Footer("下面會出現的東西","https://discordapp.com/api/webhooks/710112845567623238/_LxpGhvlK6Wp-LgavxRV1UlcsEdqnnznHXuA6d4v7YWUBCfHOBRPwCo2B7-ao0_3EbwV"),
                new Embedobj.Thumbnail("https://img.icons8.com/color/48/000000/drupal.png"),
                new Embedobj.Image("https://i.imgur.com/zul9zzl.jpg"),
                new Embedobj.Author("匿名機器人v0.0.2 訊息就放靠北用的網址", "https://kryptongta.com/images/kryptontitle2.png","https://img.icons8.com/color/144/000000/drupal.png"),
                null//new Embedobj.Field("這裡誰管的叫你們老大出來","NOVA", true)
        );
    
        webhook.addEmbed(em);
    
        webhook.excute();
        
    }
   
}

