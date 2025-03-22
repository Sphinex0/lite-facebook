
let socket = null;
const ports = new Set()

self.onconnect = (event) => {
    const port = event.ports[0]
    console.log("port", port)
    ports.add(port)
    port.onmessage = (event) => {
        const { kind, payload } = event.data
        if (kind == "connect") {
            if (!socket) {
                socket = new WebSocket(`${payload}/ws`)
                socket.onopen = () => {
                    console.log("socket is open now")
                }
                socket.onmessage = (event) => {
                    const msg = JSON.parse(event.data)
                    ports.forEach((p => p.postMessage(msg)))
                }
                socket.onerror = (error) => {
                    console.error('WebSocket error:', error);
                };
                socket.onclose = () => {
                    console.log('WebSocket connection closed');
                    socket = null;
                };
            } else {
                if (socket) {
                    socket.send(JSON.stringify({ type: "conversations" }))
                }
            }
        } else if (kind == "send") {
            if (socket && socket.readyState == WebSocket.OPEN) {
                socket.send(JSON.stringify(payload))
            } else {
                console.error('WebSocket is not open');
            }
        } else if (kind == "image") {
            if (socket && socket.readyState == WebSocket.OPEN) {
                socket.send(payload)
            } else {
                console.error('WebSocket is not open');
            }
        }
    }
}