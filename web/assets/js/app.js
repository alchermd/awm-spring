const form = document.getElementById("messageForm");
form.addEventListener("submit", function(e) {
	e.preventDefault();

	const formData = new FormData(this);
	const payload = {
		name: formData.get("name"),
		message: formData.get("message"),
		reply_me: formData.get("reply_me") === "on",
	};
	const endpoint = this.getAttribute("action");

	fetch(endpoint, {
		method: "POST",
		body: JSON.stringify(payload),
		headers: {
			"Content-Type": "application/json"
		},
	})
	.then(res => res.json())
	.then(data => {
		alert("Your message has been received");
		window.location = "/";
	});
});
