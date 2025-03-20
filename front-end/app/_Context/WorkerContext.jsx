// // Context/WorkerContext.jsx
// "use client";
// import { createContext, useContext, useEffect, useRef, useState } from "react";

// const WorkerContext = createContext();

// export function WorkerProvider({ children }) {
//     const workerRef = useRef(null);
//     const portRef = useRef(null);
//     const [clientWorker, setClientWorker] = useState(null);
//     const [conversations, setConversations] = useState([]);

//     useEffect(() => {
//         workerRef.current = new SharedWorker("/sharedworker.js");
//         portRef.current = workerRef.current.port;
//         portRef.current.start();
//         setClientWorker(workerRef.current);

//         return () => {
//             portRef.current.close();
//             portRef.current = null;
//         };
//     }, []);

//     return (
//         <WorkerContext.Provider value={{ port: portRef.current , clientWorker, conversations, setConversations , portRef }}>
//             {children}
//         </WorkerContext.Provider>
//     );
// }

// export const useWorker = () => useContext(WorkerContext);



// Context/WorkerContext.jsx
"use client";
import { createContext, useContext, useEffect, useRef, useState  } from "react";

const WorkerContext = createContext();

export function WorkerProvider({ children }) {
    const workerRef = useRef(null);
    const portRef = useRef(null);
    const [clientWorker, setClientWorker] = useState(null);
    const [conversations, setConversations] = useState([]);
    const selectedConversationRef = useRef(null);
    const [messages, setMessages] = useState([]);
    const userRef = useRef(null);

    useEffect(() => {
        const storedUser = localStorage.getItem("user");
        console.log("storedUser", storedUser);
        userRef.current = storedUser ? JSON.parse(storedUser) : null;
        console.log("userRef", userRef.current);
    }, []);

    useEffect(() => {
        const worker = new SharedWorker("/sharedworker.js");
        workerRef.current = worker;
        portRef.current = worker.port;
        portRef.current.start();
        setClientWorker(worker);

        ////////////

        const port = portRef.current;
        if (!port) return;
        const messageHandler = ({ data }) => {
            switch (data.type) {
                case "conversations":
                    const onlineUsers = data.online_users;
                    setConversations(
                        data.conversations?.map((conv) => {
                            if (conv.user_info) {
                                return {
                                    ...conv,
                                    user_info: {
                                        ...conv.user_info,
                                        online: onlineUsers?.includes(conv.user_info.id),
                                    },
                                };
                            }
                            return conv;
                        })
                    );
                    break;

                case "online":
                case "offline":
                    setConversations((prev) =>
                        prev.map((conv) => {
                            if (conv.user_info?.id === data.user_info.id) {
                                return {
                                    ...conv,
                                    user_info: {
                                        ...conv.user_info,
                                        online: data.type === "online",
                                    },
                                };
                            }
                            return conv;
                        })
                    );
                    break;

                case "new_message":
                    const msg = data.message;
                    const conversationId = msg.conversation_id;

                    setConversations((prev) => {
                        const conversation = prev.find((c) => c.conversation.id === conversationId);
                        if (conversation) {
                            return [{
                                ...conversation,
                                last_message: data?.message?.content,
                                seen: conversation.conversation.id === selectedConversationRef.current?.id ? 0 : conversation.seen + 1,
                            }, ...prev.filter((c) => c.conversation.id !== conversationId)];
                        } else {
                            return [
                                {
                                    conversation: { id: msg.conversation_id },
                                    user_info: { ...data.user_info, online: true },
                                    last_message: data?.message?.content,
                                    seen: 1
                                },
                                ...prev,
                            ];
                        }
                    });

                    console.log("selectedConversationRef", selectedConversationRef.current);
                    console.log("conversationId", conversationId);
                    if (selectedConversationRef.current?.id === conversationId) {
                        setMessages((prev) => [...prev, data]);
                        console.log("userRef.current => ", userRef.current);
                        if (userRef.current && userRef.current.id !== data.user_info.id) {
                            console.log("user", user);
                            console.log("data", data);
                            const type = selectedConversationRef.current.type == "private" ? "read_messages_private" : "read_messages_group";
                            port.postMessage({
                                kind: "send",
                                payload: {
                                    type,
                                    message: {
                                        conversation_id: conversationId,
                                    },
                                },
                            });
                        }
                        // const type = selectedConversationRef.current.type == "private" ? "read_messages_private" : "read_messages_group";
                        // port.postMessage({
                        //     kind: "send",
                        //     payload: {
                        //         type,
                        //         message: {
                        //             conversation_id: conversationId,
                        //         },
                        //     },
                        // });
                    }
                    break;

                default:
                    console.warn("Unhandled message type:", data.type);
            }
        };


        port.addEventListener("message", messageHandler);
        port.postMessage({ kind: "connect" });
        // console.log("port", port);

        // return () => {
        //     port.removeEventListener("message", messageHandler);
        // };

        ////////////


        return () => {
            port.removeEventListener("message", messageHandler);
            portRef.current?.close();
            portRef.current = null;
            workerRef.current = null;
        };
    }, []);

    return (
        <WorkerContext.Provider
            value={{ portRef, clientWorker, conversations, setConversations , selectedConversationRef , messages , setMessages }}
        >
            {children}
        </WorkerContext.Provider>
    );
}

export const useWorker = () => useContext(WorkerContext);