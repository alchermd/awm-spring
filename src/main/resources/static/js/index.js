const vm = new Vue({
    el: "#app",
    data: {
        shouldShowMessageForm: false,
        message: {
            name: "",
            body: "",
            reply_me: false,
        },
    },
    methods: {
        showMessageForm() {
            this.shouldShowMessageForm = true;
        },
        closeMessageForm() {
            this.shouldShowMessageForm = false;
        },
        sendMessage() {
            const payload = JSON.stringify(this.message);
            fetch("/api/messages/", {
                headers: {
                    "Content-Type": "application/json",
                },
                method: "POST",
                body: payload,
            })
            .then(res => res.json())
            .then(data => {
                console.log(data);
                alert("Message sent!");
            });
        }
    }
});