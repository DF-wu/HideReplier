package dfder.hidereplyer.repo;



import dfder.hidereplyer.Entity.Discord.StoreData;
import org.springframework.data.mongodb.repository.MongoRepository;


// store posted post which had posted to discord

public interface DiscordPostDataRepo extends MongoRepository<StoreData,String > {
    StoreData findBySerialNumberOrderBySerialNumberAsc();
    
    
}
