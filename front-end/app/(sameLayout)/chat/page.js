"use client";
import { useEffect, useState, useRef } from "react";
import styles from "./styles.module.css";
import Message from "@/app/(sameLayout)/chat/_components/message";
import { AddPhotoAlternate, Cancel, EmojiEmotions, Send } from "@mui/icons-material";
import { emojis } from "./_components/emojis";

export default function Chat() {
    const [clientWorker, setClientWorker] = useState(null);
    const [message, setMessage] = useState("");
    const [conversations, setConversations] = useState([]);
    const [messages, setMessages] = useState([]);
    const [selectedConversation, setSelectedConversation] = useState(null);
    const [img, setImage] = useState(null)
    const [emoji, setEmoji] = useState(false)
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
        console.log("conversations => ", conversations)
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
                    const onlineUsers = data.online_users;
                    setConversations(data.conversations.map(conv => {
                        // Fixed: Handle group conversations properly
                        if (conv.user_info) {
                            return {
                                ...conv,
                                user_info: {
                                    ...conv.user_info,
                                    online: onlineUsers?.includes(conv.user_info.id)
                                }
                            };
                        }
                        return conv; // Keep group conversations as-is
                    }));
                    break;

                case "online":
                case "offline":
                    // Fixed: Assuming server sends user ID in data.message.user_id
                    setConversations(prev => prev.map(conv => {
                        if (conv.user_info?.id === data.message.user_id) {
                            return {
                                ...conv,
                                user_info: {
                                    ...conv.user_info,
                                    online: data.type === "online"
                                }
                            };
                        }
                        return conv;
                    }));
                    break;

                case "new_message":
                    // Fixed: Use data.message instead of data
                    const msg = data.message;
                    const conversationId = msg.conversation_id;

                    setConversations(prev => {
                        const conversation = prev.find(c => c.conversation.id === conversationId);
                        if (!conversation) return prev;
                        return [conversation, ...prev.filter(c => c.conversation.id !== conversationId)];
                    });

                    if (selectedConversationRef.current?.id === conversationId) {
                        setMessages(prev => [...prev, msg]); // Fixed: Use msg instead of data
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
            kind: "image",
            payload: combinadeRef.current,
        });
        CancelFile()
    }
    const CancelFile = () => {
        combinadeRef.current = null
        setImage(null)
    }

    const HandelImage = (file) => {
        setImage(URL.createObjectURL(file));
        const reader = new FileReader();

        reader.onloadend = (e) => {
            const metadata = {
                type: file.name,
                message: {
                    conversation_id: selectedConversation.id,
                },
            };

            const metadataStr = JSON.stringify(metadata);
            const encodedMetadata = new TextEncoder().encode(metadataStr);

            const metadataLength = new Uint32Array([encodedMetadata.byteLength]);

            const fileData = new Uint8Array(e.target.result);

            const totalLength =
                4 +
                encodedMetadata.byteLength +
                fileData.byteLength;

            const combined = new Uint8Array(totalLength);

            combined.set(new Uint8Array(metadataLength.buffer), 0);
            combined.set(encodedMetadata, 4);
            combined.set(fileData, 4 + encodedMetadata.byteLength);

            combinadeRef.current = combined;
        };

        reader.readAsArrayBuffer(file);
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
        (c) => c?.conversation.id === selectedConversation?.id
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

                {
                    emoji && (
                        <div className={styles.Emojis}>
                            {
                                emojis.map((emo) => {
                                    return <div
                                        className={styles.Emoji}
                                        key={emo}
                                        onClick={() => {
                                            setMessage((prev) => prev + emo)
                                        }}
                                    >
                                        {emo}
                                    </div>
                                })
                            }
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
                    <div className={`${styles.addImageInChat}`}>
                        <label htmlFor="addImageInChat" >
                            <AddPhotoAlternate />
                        </label>
                        <label className={``} onClick={() => setEmoji((prev) => !prev)}>
                            <EmojiEmotions />
                        </label>
                    </div>
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
                            <h5 className={styles.conversationTitle}>
                                {
                                    user_info.online && " online "
                                }
                                {displayText}

                            </h5>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}