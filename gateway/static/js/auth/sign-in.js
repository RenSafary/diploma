document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("login-form");
    const ws = new WebSocket("ws://localhost:8080/sign-in-ws");

    ws.onopen = () => console.log("WebSocket connected");

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);

        if (data.status) {
            alert("Вход успешен!");

            fetch("/sign-in/put-token", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ token: data.token })
            })
            .then(res => {
                if (res.ok) {
                    console.log("Token stored in cookie");
                    window.location.href = "/"; 
                } else {
                    console.error("Failed to set token");
                }
            });
        } else {
            alert("Неверный логин или пароль");
        }
    };

    form.addEventListener("submit", (e) => {
        e.preventDefault();
        ws.send(JSON.stringify({
            username: form.username.value,
            password: form.password.value
        }));
    });
});
