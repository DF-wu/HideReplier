package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.*;
import dfder.hidereplyer.repo.CounterRepo;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.awt.*;
import java.io.IOException;
import java.time.*;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

@Service
public class DiscordService {
    
  
    
    @Autowired
    private DiscordPostDataRepo repo;
    @Autowired
    private CounterRepo counterRepo;
    
    private SerialCounter sc;

    
    private static String  url = "https://discordapp.com/api/webhooks/719110538235346955/m6VbyiiJajitpt1MlL95FW3L9B3v71nqMG1_FBTEueZMEiFwXNbxJRZZWh72Z-77LCzA";
    
    @Autowired
    public DiscordService(DiscordPostDataRepo repo, CounterRepo counterRepo)
    {
        this.repo = repo;
        this.counterRepo = counterRepo;
        sc = counterRepo.findAll().get(0);
    }
    
    
    
    // post a post to discord
    public DiscordMessage postApost(RecivedJSONofDiscordMessage recivedJSONofDiscordMessage) throws IOException
    {
        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
        Instant instant = TWtime.toInstant(ZoneOffset.of("+8"));
        long TWtimeSecont = instant.getEpochSecond();
    
        //暫時想不到比較好的寫法 要讓瓶頸不會卡在DB
        
        //++流水號
        sc.plusCounter();
        
        //+流水號
        DiscordMessage dm = recivedJSONofDiscordMessage;  // 轉成discord要求的物件
        
        //get color from client
        // 把 # 吃掉
        String hexColor = (String) recivedJSONofDiscordMessage.getExtra("color");
        hexColor = hexColor.replace("#","");
        hexColor = Integer.valueOf(hexColor,16).toString();
        
        // 蓋回去 最後傳給discord的
        recivedJSONofDiscordMessage.setExtras("color",hexColor);
        
        
        //0.0.2版格式欄位
//        em.makeEmbed("機器人名字",//"不是我沒有權限刪除阿....",
//                "靠北訊息 ",
//                "https://img.icons8.com/color/288/000000/drupal.png",
//                Color.YELLOW,
//                null,//new Embedobj.Footer("下面會出現的東西","https://discordapp.com/api/webhooks/710112845567623238/_LxpGhvlK6Wp-LgavxRV1UlcsEdqnnznHXuA6d4v7YWUBCfHOBRPwCo2B7-ao0_3EbwV"),
//                new Embedobj.Thumbnail("https://img.icons8.com/color/48/000000/drupal.png"),
//                new Embedobj.Image("https://i.imgur.com/zul9zzl.jpg"),
//                new Embedobj.Author("匿名機器人v0.0.2 訊息就放靠北用的網址", "https://kryptongta.com/images/kryptontitle2.png","https://img.icons8.com/color/144/000000/drupal.png"),
//                null//new Embedobj.Field("這裡誰管的叫你們老大出來","NOVA", true)
//        );
        dm.addEmbed(new Embedobj().makeEmbed(
                recivedJSONofDiscordMessage.getUsername(),
                recivedJSONofDiscordMessage.getContent(),
                "https://img.icons8.com/color/48/000000/drupal.png", // this post link
                hexColor,
                null, // footer
                new Embedobj.Thumbnail("https://img.icons8.com/color/48/000000/drupal.png"), //機器人縮圖
                new Embedobj.Image((String) recivedJSONofDiscordMessage.getExtra("imgUrl")),  //上傳圖片連結
                new Embedobj.Author("匿名機器人v0.0.2", "https://kryptongta.com/images/kryptontitle2.png","https://img.icons8.com/color/144/000000/drupal.png"),
                    null
                ));
        
        
        DiscordStoreData discordStoreData =
                new DiscordStoreData(TWtimeSecont, sc.getCounter(), dm);
        
        
        //清掉discord內文的部分
        recivedJSONofDiscordMessage.setContent("");
        
        
        repo.insert(discordStoreData);
        recivedJSONofDiscordMessage.excute();
        
        //流水號存回去
        counterRepo.save(sc);
        return recivedJSONofDiscordMessage;
    }
    
   
    
    public ArrayList<DiscordStoreData> gethistorylist(){
        ArrayList<DiscordStoreData> arr = (ArrayList<DiscordStoreData>)repo.findAll();
        arr.sort(Comparator.comparingInt(DiscordStoreData::getSerialNumber));
        return arr;
        
    }
    
//
//    public DiscordMessage initTheWorld(RecivedJSONofDiscordMessage recivedJSONofDiscordMessage) throws IOException
//    {
//        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
//        DiscordMessage dm = new DiscordMessage(DiscordMessage.defaultUrl);
//        Color jsonColor = Color.decode((String) recivedJSONofDiscordMessage.getExtra("color"));
//        DiscordStoreData discordStoreData =
//                new DiscordStoreData(TWtime, 1, dm, jsonColor);
//        repo.insert(discordStoreData);
//        dm.excute();
//        SerialCounter serialCounter = new SerialCounter();
//        serialCounter.setCounter(1);
//        counterRepo.save(serialCounter);
//        return dm;
//    }
//
//
    
    
    // -----------below for testing--------------
    
    
    public void testService() throws IOException
    {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }
    
    
    
    
    
 
}
