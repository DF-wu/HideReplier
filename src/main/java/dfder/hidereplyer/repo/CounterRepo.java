package dfder.hidereplyer.repo;

import dfder.hidereplyer.Entity.SerialCounter;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface CounterRepo extends MongoRepository<SerialCounter,String> {
}
