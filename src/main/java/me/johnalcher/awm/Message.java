package me.johnalcher.awm;

import lombok.Data;

@Data
public class Message {
    private Long id;
    private String body;
    private String name;
    private Boolean replyMe;

    public Message(Long id, String body, String name, Boolean replyMe) {
        this.id = id;
        this.body = body;
        this.name = name;
        this.replyMe = replyMe;
    }
}
