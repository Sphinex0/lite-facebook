"use client";
import { useEffect, useState, useRef } from "react";
import styles from "./styles.module.css";
import Message from "@/app/(sameLayout)/chat/_components/message";
import { AddPhotoAlternate, Cancel, Send } from "@mui/icons-material";
import Image from "next/image";

export default function Chat() {
    const [clientWorker, setClientWorker] = useState(null);
    const [message, setMessage] = useState("");
    const [conversations, setConversations] = useState([]);
    const [messages, setMessages] = useState([]);
    const [selectedConversation, setSelectedConversation] = useState(null);
    const [img, setImage] = useState(null)
    const selectedConversationRef = useRef(selectedConversation);
    const workerPortRef = useRef(null);
    const chatEndRef = useRef(null);
    const beforeRef = useRef(Math.floor(new Date().getTime() / 1000));
    const conversationsRef = useRef(conversations)
    const combinadeRef = useRef(null)

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

    useEffect(() => {
        conversationsRef.current = conversations
    }, [conversations])

    useEffect(() => {
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
                setMessages(data);
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
                    const conversationId = data.message.conversation_id;
                    const topCnv = conversationsRef.current.find(cnv => cnv.conversation.id === conversationId);

                    if (topCnv) {
                        const newConversations = [
                            topCnv,
                            ...conversationsRef.current.filter(cnv => cnv.conversation.id !== conversationId)
                        ];
                        setConversations(newConversations);
                    }
                    if (selectedConversationRef.current?.id === conversationId) {
                        setMessages((prev) => [...prev, data]); // Use data.message if API differs
                    } else {
                        alert("message");
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

    const SendFile = () => {
        combinadeRef && workerPortRef.current.postMessage({
            kind: "send",
            payload: combinadeRef,
        });
        CancelFile()
    }
    const CancelFile = () => {
        combinadeRef.current = null
        setImage(null)
    }

    const HandelImage = (file) => {
        setImage(URL.createObjectURL(file));
        const reader = new FileReader()
        const onloadFile = (e) => {

            const metadata = {
                type: "image",
                message: {
                    conversation_id: selectedConversation.id,
                },
            }
            const metadataStr = JSON.stringify(metadata)
            const methaDataLen = new Uint8Array([metadataStr.length])
            const fileData = e.target.result
            const totalLenght = 4 + metadataStr.length + fileData.byteLength
            const combinade = new Uint8Array(totalLenght)
            combinade.set(new Uint8Array(methaDataLen.buffer), 0)
            combinade.set(new Uint8Array(new TextEncoder().encode(metadataStr)), 4)
            combinade.set(new Uint8Array(fileData), metadataStr.length + 4)
            console.log(combinade)
            combinadeRef.current = combinade
        }
        reader.onloadend = onloadFile
        reader.readAsArrayBuffer(file)

    }

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

                {
                    img && (
                        <div className={styles.Parent}>
                            <img className={styles.imagePreview} src={img} alt="image" />
                            <div>
                                <button onClick={CancelFile} > <Cancel /> </button>
                                <button onClick={SendFile} > <Send /> </button>
                            </div>
                        </div>
                    )
                }

                <div className={styles.groupInputs}>
                    <input
                        className={styles.chatInput}
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        onKeyDown={handleSendMessage}
                        placeholder="Type your message..."
                        disabled={!selectedConversation}
                    />
                    <label htmlFor="addImageInChat" className={`${styles.addImageInChat}`}>
                        <AddPhotoAlternate />
                    </label>
                    <input
                        disabled={!selectedConversation}
                        id="addImageInChat"
                        className={styles.inputFile}
                        onChange={(event) => HandelImage(event.target.files[0])}
                        type="file"
                    />
                </div>
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