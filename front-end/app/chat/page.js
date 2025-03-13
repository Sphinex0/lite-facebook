"use client";
import { useEffect, useState, useRef } from "react";
import styles from "./styles.module.css";
import Message from "@/app/chat/_components/message";

export default function Chat() {
    const [clientWorker, setClientWorker] = useState(null);
    const [message, setMessage] = useState("");
    const [conversations, setConversations] = useState([]);
    const [messages, setMessages] = useState([]);
    const [selectedConversation, setSelectedConversation] = useState(null);
    const selectedConversationRef = useRef(selectedConversation);
    const workerPortRef = useRef(null);
    const chatEndRef = useRef(null);
    const beforeRef = useRef(Math.floor(new Date().getTime() / 1000));

    // Initialize SharedWorker
    useEffect(() => {
        const worker = new SharedWorker("/sharedworker.js");
        workerPortRef.current = worker.port;
        setClientWorker(worker);

        return () => {
            worker.port.close();
            workerPortRef.current = null;
        };
    }, []);

    // Update selectedConversationRef and fetch message history
    useEffect(() => {
        console.log("hahna")
        selectedConversationRef.current = selectedConversation;
        const fetchMessages = async () => {
            if (!selectedConversation) return;
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/messageshestories`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ conversation_id: selectedConversation.id, before: beforeRef.current }),
                credentials: "include",
            });
            if (res.ok) {
                const data = await res.json() || [];
                setMessages(data); // Set messages directly since we reset to [] on selection
            } else {
                console.error("Error fetching messages");
            }
        };
        fetchMessages();
    }, [selectedConversation]);

    // Setup message handler for SharedWorker
    useEffect(() => {
        if (!workerPortRef.current) return;

        const port = workerPortRef.current;
        port.start();

        const messageHandler = ({ data }) => {
            switch (data.type) {
                case "conversations":
                    setConversations(data.conversations);
                    break;
                case "new_message":
                    if (selectedConversationRef.current?.id === data.message.conversation_id) {
                        setMessages((prev) => [...prev, data]);
                    } else {
                        alert("mssage")
                    }
                    break;
                default:
                    console.warn("Unhandled message type:", data.type);
            }
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
    }, [messages]);

    const handleSendMessage = (event) => {
        if (event.key !== "Enter" || !message.trim()) return;
        workerPortRef.current.postMessage({
            kind: "send",
            payload: {
                type: "new_message",
                message: {
                    conversation_id: selectedConversation.id,
                    content: message,
                },
            },
        });
        setMessage("");
    };

    const handleSetSelectedConversation = (conversation) => {
        if (selectedConversation?.id != conversation.id) {
            setMessages([]);
            beforeRef.current = Math.floor(new Date().getTime() / 1000)
            setSelectedConversation(conversation);
        }
    };

    // Compute display title for the selected conversation
    const selectedConversationInfo = conversations.find(
        (c) => c.conversation.id === selectedConversation?.id
    );
    const displayTitle = selectedConversationInfo
        ? selectedConversationInfo.group?.title ||
        `${selectedConversationInfo.user_info?.first_name} ${selectedConversationInfo.user_info?.last_name}`
        : "Select a conversation";

    return (
        <div className={styles.container}>
            <div className={styles.chatContainer}>
                <div className={styles.chatHeader}>
                    <h4>{displayTitle}</h4>
                </div>

                <div className={styles.chatBody}>
                    {selectedConversation ? (
                        messages.length > 0 ? (
                            <>
                                {messages.map((msg) => {
                                    return <Message msg={msg} key={msg.message.id} />
                                })}
                            </>
                        ) : (
                            <div className={styles.emptyState}>
                                No messages in this conversation
                            </div>
                        )
                    ) : (
                        <div className={styles.emptyState}>
                            Please select a conversation
                        </div>
                    )}
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
                    const { conversation, user_info, group } = conversationInfo;
                    const displayText =
                        group?.title || `${user_info?.first_name} ${user_info?.last_name}`;
                    return (
                        <div
                            key={`conv-${conversation.id}`}
                            className={`${styles.conversationItem} ${selectedConversation?.id === conversation.id ? styles.active : ""
                                }`}
                            onClick={() => handleSetSelectedConversation(conversation)}
                        >
                            <h5 className={styles.conversationTitle}>{displayText}</h5>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}