<<<<<<< HEAD

let socket = null;
const ports = new Set()

self.onconnect = (event) => {
    // console.log(event)
    const port = event.ports[0]
    ports.add(port)
    // console.log(ports)   
    port.onmessage = (event) => {
        // console.log(event)
        const { data } = event
        if (data.kind == "connect") {
            if (!socket) {
                socket = new WebSocket("http://localhost:8080/ws")
                console.log(socket)
                socket.onopen = () => {
                    console.log("socket is open now")
                }
                socket.onmessage = (event) => {
                    console.log("event => ", event)
                    const msg = JSON.parse(event.data)
                    console.log("msg => ", msg)
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
                // console.log(socket)
                // if (socket) {
                //     socket.send({ type: "conversations" })
                // }
            }
        } else if (data.kind == "sent") {
            console.log("kkkkkkkkkkk")
            console.log(data.payload)
        }
    }
=======

let socket = null;
const ports = new Set()

self.onconnect = (event) => {
    // console.log(event)
    const port = event.ports[0]
    ports.add(port)
    // console.log(ports)   
    port.onmessage = (event) => {
        // console.log(event)
        const { kind, payload } = event.data
        if (kind == "connect") {
            if (!socket) {
                socket = new WebSocket("http://localhost:8080/ws")
                console.log(socket)
                socket.onopen = () => {
                    console.log("socket is open now")
                }
                socket.onmessage = (event) => {
                    console.log("event => ", event)
                    const msg = JSON.parse(event.data)
                    console.log("msg => ", msg)
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
                // console.log(socket)
                // if (socket) {
                //     socket.send({ type: "conversations" })
                // }
            }
        } else if (kind == "send") {
            console.log("payload => ", payload)
            if (socket && socket.readyState == WebSocket.OPEN) {
                console.log("soket hya ", socket)
                socket.send(JSON.stringify(payload))
            } else {
                console.error('WebSocket is not open');
            }
        } else if (kind == "image") {
            if (socket && socket.readyState == WebSocket.OPEN) {
                console.log("soket hya ", socket)
                console.log("payload", payload)
                socket.send(payload)
            } else {
                console.error('WebSocket is not open');
            }
        }
    }
>>>>>>> main
}