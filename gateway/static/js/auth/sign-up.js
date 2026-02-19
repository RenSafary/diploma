document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("sign-up-form");

    form.addEventListener("submit", (e) => {
        e.preventDefault();

        const ws = new WebSocket("ws://localhost:8080/sign-up-ws");

        ws.onopen = () => {
            ws.send(JSON.stringify({
                username: form.username.value,
                password: form.password.value,
                firstname: form.firstname.value,
                lastname: form.lastname.value,
                email: form.email.value,
                age: form.age.value,
                sex: form.sex.value
            }));
        };

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);

            if (data.status) {
                alert("Регистрация успешна!");

                fetch("/put-token", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ token: data.token })
                })
                .then(res => {
                    if (res.ok) window.location.href = "/";
                });

            } else {
                alert(data.error || "Ошибка регистрации");
            }

            ws.close();
        };

        ws.onerror = () => {
            alert("Ошибка соединения с сервером");
        };
    });
});
