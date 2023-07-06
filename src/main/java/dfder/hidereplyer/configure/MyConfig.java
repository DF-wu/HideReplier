package dfder.hidereplyer.configure;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConfigurationProperties(prefix = "hidereplier")
public class MyConfig {
    private static String version;
    private static String hostURL;

    private static String DCWebhook;

    public static String getDCWebhook() {
        return DCWebhook;
    }

    public static void setDCWebhook(String DCWebhook) {
        MyConfig.DCWebhook = DCWebhook;
    }

    public String getHostURL() {
        return hostURL;
    }

    public void setHostURL(String hostURL) {
        MyConfig.hostURL = hostURL;
    }
    //-------------getter & setter start-----------
    
    public String getVersion()
    {
        return version;
    }
    
    public void setVersion(String version)
    {
        this.version = version;
    }

    
    //-------------getter & setter end-----------
}