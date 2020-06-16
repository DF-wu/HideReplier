package dfder.hidereplyer.Entity;

import java.awt.*;
import java.util.ArrayList;
import java.util.List;

public class Embedobj {
    
    private String title;
    private String description;
    private String url;
    private int color;
    
    private Footer footer;
    private Thumbnail thumbnail;
    private Image image;
    private Author author;
    
    private ArrayList<Field> fields = new ArrayList<>();
    
    
    public ArrayList<Field> getFields(){
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
    
    public int getColor()
    {
        return color;
    }
    
    public void setColor(Color color)
    {
        this.color = color.getRGB();
    }
    
    
    public static class Footer {
        private String text;
        private String iconUrl;
        
        public Footer(String text, String iconUrl) {
            this.text = text;
            this.iconUrl = iconUrl;
        }
    }
    
    public static class Thumbnail {
        private String url;
        
        public Thumbnail(String url) {
            this.url = url;
        }
    }
    
    public static class Image {
        private String url;
        
        public Image(String url) {
            this.url = url;
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
        
    }
    
    
    public Embedobj makeEmbed(String title, String description, String url, Color color, Footer footer, Thumbnail thumbnail, Image image, Author author , Field field ){
        this.footer = footer;
        this.thumbnail = thumbnail;
        this.title = title;
        this.description= description;
        this.url = url;
        this.color = getDecimalColor(color);
        this.image=image;
        this.author = author;
        this.fields.add(field);
        return this;
    }
    
    public int getDecimalColor(Color color){
        int rgb = color.getRed();
        rgb = (rgb << 8) + color.getGreen();
        rgb = (rgb << 8) + color.getBlue();
        return rgb;
    }
}
