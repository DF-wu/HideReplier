package dfder.hidereplyer.Entity.Discord;

import com.google.gson.Gson;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Field;

import javax.net.ssl.HttpsURLConnection;
import java.io.IOException;
import java.io.OutputStream;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;


public class PostMessage {
    @Id
    private String id;
    
    public final static String defaultUrl = System.getenv("DC_WEBHOOK_URL");

    
    
    @Field
    private String url;
    @Field
    private String content;
    @Field
    private String username;
    @Field
    private String avatar_url;
    @Field
    private boolean tts;
    @Field
    private ArrayList<Embedobj> embeds = new ArrayList<>();
    
    public PostMessage()
    {
    }
    
    
    public String getUrl()
    {
        return url;
    }
    
    public void setUrl(String url)
    {
        this.url = url;
    }
    
    
    public void excute() throws IOException
    {
        Gson gson = new Gson();
        String json = gson.toJson(this);
        System.out.print(json);
    
    
        URL url = new URL(this.url);
        HttpsURLConnection connection = (HttpsURLConnection) url.openConnection();
        connection.setRequestMethod("POST");
        connection.addRequestProperty("Content-Type", "application/json");
        connection.addRequestProperty("User-Agent", "D.F./1.0");
    
    
        connection.setDoOutput(true);
        OutputStream stream = connection.getOutputStream();
        stream.write(json.getBytes(StandardCharsets.UTF_8));
        //original : stream.write(json.toString().getBytes());
        //change: covert to UTF-8
        stream.flush();
        stream.close();
    
        // DF add
        // response
        int responseCode = connection.getResponseCode();
        System.out.println("\nSending 'POST' request to URL : " + url);
        //System.out.println("Post body " + json);
        System.out.println("Response Code : " + responseCode);

        if (responseCode < 200 || responseCode >= 300) {
            throw new IOException("Discord webhook request failed with status " + responseCode);
        }
        
        
//        BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
//        String inputLine;
//        StringBuilder response = new StringBuilder();
//
//        while ((inputLine = in.readLine()) != null) {
//            response.append(inputLine);
//        }
//        System.out.println("!! response :" + response);
        //df add end
    
        // default
        //connection.getInputStream().close(); //I'm not sure why but it doesn't work without getting the InputStream
        connection.disconnect();
        
    }
    
    
    
    
    public void addEmbed(Embedobj embedobj){
        this.embeds.add(embedobj);
    }
    
    public PostMessage(String url)
    {
        if (url == null || url.isBlank()) {
            throw new IllegalArgumentException("Discord webhook URL is required");
        }
        this.url = url;
    }
    
    public String getContent()
    {
        return content;
    }
    
    public void setContent(String content)
    {
        this.content = content;
    }
    
    public String getUsername()
    {
        return username;
    }
    
    public void setUsername(String username)
    {
        this.username = username;
    }
    
    public String getAvatarUrl()
    {
        return avatar_url;
    }
    
    public void setAvatarUrl(String avatarUrl)
    {
        this.avatar_url = avatarUrl;
    }
    
    public boolean isTts()
    {
        return tts;
    }
    
    public void setTts(boolean tts)
    {
        this.tts = tts;
    }
    

}
