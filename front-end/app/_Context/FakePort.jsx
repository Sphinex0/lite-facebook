
export default class FakePort {
    constructor(url) {
        this.listeners = [];
        this.socket = new WebSocket(`${url}/api/ws`);
        this.socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            this.listeners.forEach(listener => listener({ data }));
        };
        this.socket.onopen = () => console.log("open")
        this.socket.onerror = (error) => console.log("WebSocket error:", error);
        this.socket.onclose = () => console.log("WebSocket closed");
    }

    postMessage(message) {
        if (this.socket && this.socket.readyState == WebSocket.OPEN) {
            const { kind, payload } = message
            if (kind == "send") {
                this.socket.send(JSON.stringify(payload));
            } else if (kind == "image") {
                this.socket.send(payload)
            } else if (kind == "close") {
                this.socket.close()
            }
            this.socket.send(JSON.stringify(message));
        }
    }

    addEventListener(type, listener) {
        if (type === "message") {
            this.listeners.push(listener);
        }
    }

    removeEventListener(type, listener) {
        if (type === "message") {
            this.listeners = this.listeners.filter(l => l !== listener);
        }
    }

    start() {
        // No-op
    }

    close() {
        this.socket.close();
    }
}
