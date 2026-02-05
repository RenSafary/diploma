function show(id) {
    document.querySelectorAll(".content > div").forEach(el => {
        el.classList.add("hidden");
    });

    document.getElementById(id).classList.remove("hidden");
}

/* ===== WEBSOCKET ===== */

const ws = new WebSocket("ws://localhost:8080/ws/admin");

function send(action, id) {
    ws.send(JSON.stringify({
        action: action,
        id: id
    }));
}

function makeAdmin() {
    send("admin", adminId.value);
}

function makeWorker() {
    send("worker", workerId.value);
}
