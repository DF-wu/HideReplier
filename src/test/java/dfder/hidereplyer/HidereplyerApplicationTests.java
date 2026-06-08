package dfder.hidereplyer;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.ActiveProfiles;

import dfder.hidereplyer.service.DiscordService;

@SpringBootTest
@ActiveProfiles("test")
class HidereplyerApplicationTests {
    @MockBean
    private DiscordService discordService;
    
    @Test
    void contextLoads()
    {
    }
    
}
