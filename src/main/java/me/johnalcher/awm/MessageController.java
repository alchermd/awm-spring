package me.johnalcher.awm;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.atomic.AtomicLong;

@RestController
public class MessageController {
    private final AtomicLong id = new AtomicLong();

    @PostMapping("/api/messages/")
    public Message newMessage(@RequestBody Message message) {
        return new Message(id.incrementAndGet(), message.getName(), message.getBody(), message.getReplyMe());
    }
}
