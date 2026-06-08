package dfder.hidereplyer.service;


import dfder.hidereplyer.Entity.*;
import dfder.hidereplyer.Entity.Discord.PostMessage;
import dfder.hidereplyer.Entity.Discord.StoreData;
import dfder.hidereplyer.Entity.Discord.Embedobj;
import dfder.hidereplyer.configure.MyConfig;
import dfder.hidereplyer.repo.CounterRepo;
import dfder.hidereplyer.repo.DiscordPostDataRepo;
import dfder.hidereplyer.test.TestDiscordPost;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.*;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

@Service
public class DiscordService {

    private static final ZoneId TAIWAN_ZONE = ZoneId.of("+8");
    private static final ZoneOffset TAIWAN_OFFSET = ZoneOffset.of("+8");
    private static final String AUTHOR_ICON_URL = "https://img.icons8.com/color/144/000000/drupal.png";
    private static final String COLOR_KEY = "color";
    private static final String AVATAR_URL_KEY = "avatar_url";
    private static final String IP_KEY = "ip";
    private static final String THUMBNAIL_KEY = "thumbnail";
    private static final String IMAGE_URL_KEY = "imgUrl";

    private final DiscordPostDataRepo repo;
    private final CounterRepo counterRepo;
    private final MyConfig myConfig;

    private final SerialCounter serialCounter;

    @Autowired
    public DiscordService(DiscordPostDataRepo repo, CounterRepo counterRepo, MyConfig myConfig) {
        this.repo = repo;
        this.counterRepo = counterRepo;
        this.myConfig = myConfig;
        this.serialCounter = loadOrCreateCounter();
    }

    public String getVersion() {
        return myConfig.getVersion();
    }

    public PostMessage postAnonymousMessage(RecivedJSONofPostMessage postMessage) throws IOException {
        long taiwanEpochSecond = currentTaiwanEpochSecond();
        String originalContent = postMessage.getContent();
        String posterIp = readExtra(postMessage, IP_KEY);
        String normalizedColor = normalizeColor(postMessage);

        serialCounter.plusCounter();
        applyClientOverrides(postMessage);
        postMessage.addEmbed(buildEmbed(postMessage, originalContent, posterIp, normalizedColor, serialCounter.getCounter()));

        dispatchToDiscord(postMessage, originalContent);
        repo.insert(new StoreData(taiwanEpochSecond, serialCounter.getCounter(), posterIp, postMessage));
        counterRepo.save(serialCounter);
        return postMessage;
    }

    public ArrayList<StoreData> getHistoryList() {
        ArrayList<StoreData> history = new ArrayList<>(repo.findAll());
        history.sort(Comparator.comparingInt(StoreData::getSerialNumber));
        return history;
    }

    public void testService() throws IOException {
        TestDiscordPost t = new TestDiscordPost();
        t.postTest();
    }

    private SerialCounter loadOrCreateCounter() {
        List<SerialCounter> counters = counterRepo.findAll();
        if (!counters.isEmpty()) {
            return counters.get(0);
        }

        SerialCounter newCounter = new SerialCounter();
        newCounter.setCounter(0);
        return counterRepo.save(newCounter);
    }

    private long currentTaiwanEpochSecond() {
        LocalDateTime taiwanTime = LocalDateTime.now(Clock.system(TAIWAN_ZONE));
        Instant instant = taiwanTime.toInstant(TAIWAN_OFFSET);
        return instant.getEpochSecond();
    }

    private void applyClientOverrides(RecivedJSONofPostMessage postMessage) {
        postMessage.setAvatarUrl(readExtra(postMessage, AVATAR_URL_KEY));
    }

    private String normalizeColor(RecivedJSONofPostMessage postMessage) {
        String color = readExtra(postMessage, COLOR_KEY);
        String normalizedColor = Integer.valueOf(color.replace("#", ""), 16).toString();
        postMessage.setExtras(COLOR_KEY, normalizedColor);
        return normalizedColor;
    }

    private Embedobj buildEmbed(RecivedJSONofPostMessage postMessage,
                                String content,
                                String posterIp,
                                String color,
                                int serialNumber) {
        String botAuthor = "匿名機器人v" + myConfig.getVersion() + "（點我去發文）";
        Embedobj embed = new Embedobj();
        embed.makeEmbed(
                postMessage.getUsername(),
                content,
                myConfig.getHostUrl(),
                color,
                null,
                new Embedobj.Thumbnail(readExtra(postMessage, THUMBNAIL_KEY)),
                new Embedobj.Image(readExtra(postMessage, IMAGE_URL_KEY)),
                new Embedobj.Author(botAuthor, myConfig.getHostUrl(), AUTHOR_ICON_URL),
                new Embedobj.Field("流水號", String.valueOf(serialNumber), true)
        );
        embed.addField("來自：", posterIp, true);
        return embed;
    }

    private void dispatchToDiscord(RecivedJSONofPostMessage postMessage, String originalContent) throws IOException {
        postMessage.setUrl(myConfig.getDcWebhook());
        postMessage.setContent("");
        postMessage.excute();
        postMessage.setContent(originalContent);
    }

    private String readExtra(RecivedJSONofPostMessage postMessage, String key) {
        return (String) postMessage.getExtra(key);
    }
}
