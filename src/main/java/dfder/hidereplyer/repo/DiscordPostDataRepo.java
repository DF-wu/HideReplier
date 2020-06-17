package dfder.hidereplyer.repo;



import dfder.hidereplyer.Entity.DiscordStoreData;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.ArrayList;


// store posted post which had posted to discord

public interface DiscordPostDataRepo extends MongoRepository<DiscordStoreData,String > {
    DiscordStoreData findBySerialNumberOrderBySerialNumberAsc();
}
