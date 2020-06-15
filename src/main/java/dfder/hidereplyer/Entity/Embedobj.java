package dfder.hidereplyer.Entity;

import javax.crypto.ExemptionMechanism;
import java.awt.*;
import java.util.ArrayList;
import java.util.EnumMap;
import java.util.List;

public class Embedobj {
    
    private String title;
    private String description;
    private String url;
    private Color color;
    
    private Footer footer;
    private Thumbnail thumbnail;
    private Image image;
    private Author author;
    
    private List<Field> fields = new ArrayList<>();
    
    public List<Field> getFields(){
        return fields;
    }
    
    public void addField(String name, String value, Boolean inLine){
        this.fields.add(new Field(name,value,inLine));
    }
    
    public String getTitle()
    {
        return title;
    }
    
    public void setTitle(String title)
    {
        this.title = title;
    }
    
    public String getDescription()
    {
        return description;
    }
    
    public void setDescription(String description)
    {
        this.description = description;
    }
    
    public String getUrl()
    {
        return url;
    }
    
    public void setUrl(String url)
    {
        this.url = url;
    }
    
    public Color getColor()
    {
        return color;
    }
    
    public void setColor(Color color)
    {
        this.color = color;
    }
    
    
    public static class Footer {
        private String text;
        private String iconUrl;
        
        public Footer(String text, String iconUrl) {
            this.text = text;
            this.iconUrl = iconUrl;
        }
        
        private String getText() {
            return text;
        }
        
        private String getIconUrl() {
            return iconUrl;
        }
    }
    
    public static class Thumbnail {
        private String url;
        
        public Thumbnail(String url) {
            this.url = url;
        }
        
        private String getUrl() {
            return url;
        }
    }
    
    public static class Image {
        private String url;
        
        public Image(String url) {
            this.url = url;
        }
        
        private String getUrl() {
            return url;
        }
    }
    
    public static class Author {
        private String name;
        private String url;
        private String iconUrl;
        
        public Author(String name, String url, String iconUrl) {
            this.name = name;
            this.url = url;
            this.iconUrl = iconUrl;
        }
        
        private String getName() {
            return name;
        }
        
        private String getUrl() {
            return url;
        }
        
        private String getIconUrl() {
            return iconUrl;
        }
    }
    
    
    public static class Field{
        private String name;
        private String value;
        private boolean inline;
        
        public Field(String name, String value, boolean inline)
        {
            this.name = name;
            this.value = value;
            this.inline = inline;
        }
        
        public String getName()
        {
            return name;
        }
        
        public String getValue()
        {
            return value;
        }
        
        public boolean isInline()
        {
            return inline;
        }
        
    }
    
    
    public static Embedobj makeEmbed(String title, String description, String url, Color color, Footer footer, Thumbnail thumbnail, Image image, Author author , Field field ){
        Embedobj em = new Embedobj();
        em.footer = footer;
        em.thumbnail = thumbnail;
        em.title = title;
        em.description= description;
        em.url = url;
        em.color = color;
        em.footer = footer;
        em.image=image;
        em.author = author;
        em.fields.add(field);
        return em;
    }
}
