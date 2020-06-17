package dfder.hidereplyer.Entity;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

@Document(collection = "Counter")
public class SerialCounter {
    @Id
    private String id;
    @Field
    private int counter;

    public int getCounter()
    {
        return counter;
    }
    
    public void setCounter(int counter)
    {
        this.counter = counter;
    }
    
    public void plusCounter() { ++counter; }
}
