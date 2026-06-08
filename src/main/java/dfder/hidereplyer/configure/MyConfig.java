package dfder.hidereplyer.configure;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConfigurationProperties(prefix = "hidereplier")
public class MyConfig {

    private String version;
    private String hostUrl;
    private String dcWebhook;

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public String getHostUrl() {
        return hostUrl;
    }

    public void setHostUrl(String hostUrl) {
        this.hostUrl = hostUrl;
    }

    public String getDcWebhook() {
        return dcWebhook;
    }

    public void setDcWebhook(String dcWebhook) {
        this.dcWebhook = dcWebhook;
    }
}
