package dfder.hidereplyer.repo;


import dfder.hidereplyer.Entity.DiscordWebhook;
import org.springframework.data.mongodb.repository.MongoRepository;


// store posted post which had posted to discord

public interface DiscordPostDataRepo extends MongoRepository<DiscordWebhook,String > {


}
