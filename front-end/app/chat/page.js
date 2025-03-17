"use client";
import { useEffect, useState, useRef } from "react";
import styles from "./styles.module.css";
import Message from "@/app/chat/_components/message";

export default function Chat() {
    const [clientWorker, setClientWorker] = useState(null);
    const [message, setMessage] = useState("");
    const [conversations, setConversations] = useState([]);
    const [selectedConversation, setSelectedConversation] = useState(null);
    const workerPortRef = useRef(null);
    const chatEndRef = useRef(null);

    // Initialize 
    useEffect(() => {
        const worker = new SharedWorker("/sharedworker.js");
        workerPortRef.current = worker.port;
        setClientWorker(worker);

        return () => {
            worker.port.close();
            workerPortRef.current = null;
        };
    }, []);

    // Setup message handler
    useEffect(() => {
        if (!workerPortRef.current) return;

        const port = workerPortRef.current;
        port.start();

        const messageHandler = ({ data }) => {
            console.log("Received data:", data);
            if (data.type === "conversations") {
                setConversations(data.conversations);
            }
            // Add other message types here
        };

        port.addEventListener("message", messageHandler);
        port.postMessage({ kind: "connect" });

        return () => {
            port.removeEventListener("message", messageHandler);
        };
    }, []);

    // Scroll to bottom when messages change
    useEffect(() => {
        chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [conversations]);

    const handleSendMessage = (event) => {
        if (event.key !== "Enter" || !message.trim()) return;
        
        workerPortRef.current?.postMessage({
            kind: "send",
            payload: {
                type : "new_message",
                message : {
                    conversation_id : selectedConversation,
                    content : message
                }
            }
        });

        setMessage("");
    };

    const handleSelectConversation = (conversation) => {
        setSelectedConversation(conversation);
    };

    return (
        <div className={styles.container}>
            <div className={styles.chatContainer}>
                <div className={styles.chatHeader}>
                    <h4>{selectedConversation?.title || "Select a conversation"}</h4>
                </div>

                <div className={styles.chatBody} >
                    {
                        selectedConversation?.messages?.map((msg, index) => (
                            // <div key={`msg-${index}`} className={styles.message}>
                            //     <div className={styles.messageHeader}>
                            //         <span className={styles.userName}>
                            //             {msg.senderName}
                            //         </span>
                            //     </div>
                            //     <div className={styles.messageContent}>
                            //         {msg.content}
                            //     </div>
                            // </div>
                            <Message msg={msg} index={index} />
                        ))
                    }
                    <div ref={chatEndRef} />
                </div>

                <input
                    className={styles.chatInput}
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    onKeyDown={handleSendMessage}
                    placeholder="Type your message..."
                    disabled={!selectedConversation}
                />
            </div>

            <div className={styles.conversationsList}>
                {conversations.map((conversationInfo) => {
                    const { conversation, user_info, group } = conversationInfo
                    const { id } = conversation
                    const displayText = group.title || `${user_info.first_name} ${user_info.last_name}`;
                    return (
                        <div
                            key={`conv-${id}`}
                            className={`${styles.conversationItem} ${selectedConversation?.id === id ? styles.active : ""
                                }`}
                            onClick={() => handleSelectConversation(conversation)}
                        >
                            <h5 className={styles.conversationTitle}>{displayText}</h5>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}