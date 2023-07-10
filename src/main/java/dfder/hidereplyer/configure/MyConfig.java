package dfder.hidereplyer.configure;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
//@ConfigurationProperties(prefix = "hidereplier")
public class MyConfig {

    @Value("${hidereplier.version}")
    private String version;

    @Value("${hidereplier.hostURL}")
    private String hostURL;

    @Value("${hidereplier.DCWebhook}")
    private String DCWebhook;

    @Value("${TENOR_API_KEY}")
    private String TENOR_API_KEY;


    public String getVersion() {
        return version;
    }

    public String getHostURL() {
        return hostURL;
    }

    public String getDCWebhook() {
        return DCWebhook;
    }

    public String getTENOR_API_KEY() {
        return TENOR_API_KEY;
    }

//-------------getter & setter end-----------
}