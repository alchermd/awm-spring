package me.johnalcher.awm;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Date;

@RestController
public class MessageController {
    private MessageRepository repository;

    public MessageController(MessageRepository repository) {
        this.repository = repository;
    }

    @PostMapping("/api/messages/")
    public Message newMessage(@RequestBody Message message) {
        Message newMessage = new Message();

        newMessage.setName(message.getName());
        newMessage.setBody(message.getBody());
        newMessage.setReplyMe(message.getReplyMe());
        newMessage.setCreatedAt(new Date());

        return repository.save(newMessage);
    }
}
