document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("login-form");

    const ws = new WebSocket("ws://localhost:8080/sign-in-ws");

    ws.onopen = () => console.log("WebSocket connected");

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.status) {
            alert("Вход успешен! Токен: " + data.token);
            window.location.href = "/";
        } else {
            alert("Неверный логин или пароль");
        }
    };

    ws.onerror = (err) => console.error("WebSocket error:", err);

    form.addEventListener("submit", (e) => {
        e.preventDefault();
        const client = {
            username: form.username.value,
            password: form.password.value
        };
        ws.send(JSON.stringify(client));
    });
});
