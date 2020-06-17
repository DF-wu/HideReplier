package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.DiscordMessage;
import dfder.hidereplyer.Entity.DiscordStoreData;
import dfder.hidereplyer.Entity.RecivedJSONofDiscordMessage;
import dfder.hidereplyer.Entity.SerialCounter;
import dfder.hidereplyer.repo.CounterRepo;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.Clock;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.ArrayList;

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
    
        //暫時想不到比較好的寫法 要讓瓶頸不會卡在DB
        
        //++流水號
        sc.plusCounter();
        
        //+流水號
        DiscordMessage dm = recivedJSONofDiscordMessage;
        DiscordStoreData discordStoreData =
                new DiscordStoreData(TWtime, sc.getCounter(), dm);
        repo.insert(discordStoreData);
        recivedJSONofDiscordMessage.excute();
        
        //流水號存回去
        counterRepo.save(sc);
        return recivedJSONofDiscordMessage;
    }
    
   
    
    public ArrayList<DiscordStoreData> gethistorylist(){
        return (ArrayList<DiscordStoreData>) repo.findAll();
        
    }
    
    public DiscordMessage initTheWorld(RecivedJSONofDiscordMessage recivedJSONofDiscordMessage) throws IOException
    {
        final LocalDateTime TWtime = LocalDateTime.now(Clock.system(ZoneId.of("+8")));
        DiscordMessage dm = new DiscordMessage(DiscordMessage.defaultUrl);
        DiscordStoreData discordStoreData =
                new DiscordStoreData(TWtime, 1, dm);
        repo.insert(discordStoreData);
        dm.excute();
        SerialCounter serialCounter = new SerialCounter();
        serialCounter.setCounter(1);
        counterRepo.save(serialCounter);
        return dm;
    }

    
    
    
    // -----------below for testing--------------
    
    
    public void testService() throws IOException
    {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }
    
    
    
    
    
 
}
