package dfder.hidereplyer.Entity;

import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "Counter")

public class SerialCounter {
    private int counter;
    
    public int getCounter()
    {
        return counter;
    }
    
    public void setCounter(int counter)
    {
        this.counter = counter;
    }
}
