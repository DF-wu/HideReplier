package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.*;
import dfder.hidereplyer.Entity.Discord.PostMessage;
import dfder.hidereplyer.Entity.Discord.StoreData;
import dfder.hidereplyer.Entity.Discord.Embedobj;
import dfder.hidereplyer.configure.MyConfig;
import dfder.hidereplyer.repo.CounterRepo;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.*;
import java.util.ArrayList;
import java.util.Comparator;

@Service
public class DiscordService {


    private DiscordPostDataRepo repo;
    private CounterRepo counterRepo;
    private final MyConfig myConfig;
    
    private SerialCounter sc;

    @Autowired
    public DiscordService(DiscordPostDataRepo repo, CounterRepo counterRepo, MyConfig myConfig)
    {
        this.repo = repo;
        this.counterRepo = counterRepo;
        sc = counterRepo.findAll().get(0);
        this.myConfig = myConfig;
    }
    
    
    
    public String getVersion()
    {
        return myConfig.getVersion();
    }

    // post a post to discord
    public PostMessage postApost(RecivedJSONofPostMessage recivedJSONofDiscordMessage) throws IOException
    {
        // 取得台灣時間 轉成Epoch秒 後使用long儲存  以便解回
        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
        Instant instant = TWtime.toInstant(ZoneOffset.of("+8"));
        long TWtimeSecont = instant.getEpochSecond();

        //++流水號
        sc.plusCounter();

        // 轉出discord要求的物件
        PostMessage postMessage = recivedJSONofDiscordMessage;

        //hotfix: 不知為何recivedJSONofDiscordMessage的avatar_url會收不到json 暫時指定處理
        postMessage.setAvatarUrl((String) recivedJSONofDiscordMessage.getExtra("avatar_url"));

        
        
        //get color from client
        // 把 # 吃掉
        String hexColor = (String) recivedJSONofDiscordMessage.getExtra("color");
        hexColor = hexColor.replace("#","");
        hexColor = Integer.valueOf(hexColor,16).toString();

        // 蓋回去 最後要傳回給discord的
        recivedJSONofDiscordMessage.setExtras("color",hexColor);

        //把前端取得的ip取出
        String ip = (String) recivedJSONofDiscordMessage.getExtra("ip");

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

        String discordFieldcontent = recivedJSONofDiscordMessage.getContent();

        //新增Embed物件
        Embedobj em = new Embedobj();
        em.makeEmbed(
                recivedJSONofDiscordMessage.getUsername(),
                discordFieldcontent,
                myConfig.getHostURL(), // this post's link
                hexColor,
                null ,// new Embedobj.Footer("來自： " + ip,""),  // footer
                new Embedobj.Thumbnail((String) recivedJSONofDiscordMessage.getExtra("thumbnail")), //機器人縮圖
                new Embedobj.Image((String) recivedJSONofDiscordMessage.getExtra("imgUrl")),  //上傳圖片連結
                new Embedobj.Author("匿名機器人v0.4（點我去發文）", myConfig.getHostURL(), "https://img.icons8.com/color/144/000000/drupal.png"),
                new Embedobj.Field("流水號", String.valueOf(sc.getCounter()) , true)
        );
        em.addField("來自：", ip, true);
        postMessage.addEmbed(em);


        //建構儲存的Entity
        /*
        * time台灣時間
        * counter 流水號
        * ip 發文者ip
        * Discord Json 的java bean
        * */
        StoreData discordStoreData =
                new StoreData(
                        TWtimeSecont,
                        sc.getCounter(),
                        ip,
                        postMessage
                );
        //存進DB
        repo.insert(discordStoreData);




        // assign url
        recivedJSONofDiscordMessage.setUrl(myConfig.getDCWebhookURL());

        //清掉discord內文的部份 讓discord不會顯示content 而是只有embed內容
        recivedJSONofDiscordMessage.setContent("");
        // 送出post
        postMessage.excute();
        // 把content補回來 不然前端沒有
        recivedJSONofDiscordMessage.setContent(discordFieldcontent);


        //流水號存回去
        counterRepo.save(sc);
        return recivedJSONofDiscordMessage;
    }



    public ArrayList<StoreData> gethistorylist(){
        ArrayList<StoreData> arr = (ArrayList<StoreData>)repo.findAll();
        arr.sort(Comparator.comparingInt(StoreData::getSerialNumber));
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
