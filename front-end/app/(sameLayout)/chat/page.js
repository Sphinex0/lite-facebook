"use client";
import { useEffect, useState, useRef, useContext } from "react";
import styles from "./styles.module.css";
import Message from "@/app/(sameLayout)/chat/_components/message";
import { AddPhotoAlternate, Cancel, EmojiEmotions, Send } from "@mui/icons-material";
import { emojis } from "./_components/emojis";
import UserInfo from "../_components/userInfo";
import { Context } from "../layout";

export default function Chat() {
    // Destructure context, excluding conversationsRef since it's not provided
    let { clientWorker, workerPortRef, conversations, setConversations } = useContext(Context);
    console.log(workerPortRef); // Debugging workerPortRef

    // Define conversationsRef locally
    const conversationsRef = useRef(conversations);

    const [message, setMessage] = useState("");
    const [messages, setMessages] = useState([]);
    const [selectedConversation, setSelectedConversation] = useState(null);
    const [img, setImage] = useState(null);
    const [emoji, setEmoji] = useState(false);
    const selectedConversationRef = useRef(selectedConversation);
    const chatEndRef = useRef(null);
    const beforeRef = useRef(Math.floor(new Date().getTime() / 1000));
    const combinadeRef = useRef(null);

    // Update conversationsRef when conversations changes
    useEffect(() => {
        conversationsRef.current = conversations;
    }, [conversations]);

    // Fetch messages when selectedConversation changes
    useEffect(() => {
        selectedConversationRef.current = selectedConversation;
        const fetchMessages = async () => {
            if (!selectedConversation) return;
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/messageshistories`, {
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

    // Setup message handler for SharedWorker, depending on clientWorker
    useEffect(() => {
        if (!workerPortRef.current) return;
        const port = workerPortRef.current;
        port.start();

        const messageHandler = ({ data }) => {
            switch (data.type) {
                case "conversations":
                    const onlineUsers = data.online_users;
                    setConversations(
                        data.conversations.map((conv) => {
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
                            return [conversation, ...prev.filter((c) => c.conversation.id !== conversationId)];
                        } else {
                            return [
                                {
                                    conversation: { id: msg.conversation_id },
                                    user_info: { ...data.user_info, online: true },
                                },
                                ...prev,
                            ];
                        }
                    });

                    if (selectedConversationRef.current?.id === conversationId) {
                        setMessages((prev) => [...prev, data]);
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
    }, [clientWorker]); // Depend on clientWorker to re-run when it’s set

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
        if (combinadeRef.current) {
            workerPortRef.current.postMessage({
                kind: "image",
                payload: combinadeRef.current,
            });
            CancelFile();
        }
    };

    const CancelFile = () => {
        combinadeRef.current = null;
        setImage(null);
    };

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
            const totalLength = 4 + encodedMetadata.byteLength + fileData.byteLength;
            const combined = new Uint8Array(totalLength);

            combined.set(new Uint8Array(metadataLength.buffer), 0);
            combined.set(encodedMetadata, 4);
            combined.set(fileData, 4 + encodedMetadata.byteLength);

            combinadeRef.current = combined;
        };

        reader.readAsArrayBuffer(file);
    };

    const handleSetSelectedConversation = (conversation) => {
        if (selectedConversation?.id !== conversation.id) {
            setMessages([]);
            beforeRef.current = Math.floor(new Date().getTime() / 1000);
            setSelectedConversation(conversation);
        }
    };

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
                                {messages.map((msg) => (
                                    <Message msg={msg} key={msg.message.id} />
                                ))}
                            </>
                        ) : (
                            <div className={styles.emptyState}>
                                No messages in this conversation
                            </div>
                        )
                    ) : (
                        <div className={styles.emptyState}>Please select a conversation</div>
                    )}
                    <div ref={chatEndRef} />
                </div>

                {img && (
                    <div className={styles.Parent}>
                        <img className={styles.imagePreview} src={img} alt="image" />
                        <div>
                            <button onClick={CancelFile}>
                                <Cancel />
                            </button>
                            <button onClick={SendFile}>
                                <Send />
                            </button>
                        </div>
                    </div>
                )}

                {emoji && (
                    <div className={styles.Emojis}>
                        {emojis.map((emo) => (
                            <div
                                className={styles.Emoji}
                                key={emo}
                                onClick={() => setMessage((prev) => prev + emo)}
                            >
                                {emo}
                            </div>
                        ))}
                    </div>
                )}

                <div className={styles.groupInputs}>
                    <input
                        className={styles.chatInput}
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        onKeyDown={handleSendMessage}
                        placeholder="Type your message..."
                        disabled={!selectedConversation}
                    />
                    <div className={styles.addImageInChat}>
                        <label htmlFor="addImageInChat">
                            <AddPhotoAlternate />
                        </label>
                        <label onClick={() => setEmoji((prev) => !prev)}>
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
                    const { conversation, user_info, group , last_message} = conversationInfo;
                    const onlineDiv = true
                    return (
                        <div
                            key={`conv-${conversation.id}`}
                            className={`${styles.conversationItem} ${selectedConversation?.id === conversation.id ? styles.active : ""
                                }`}
                            onClick={() => handleSetSelectedConversation(conversation)}
                        >
                            <UserInfo userInfo={user_info} group={group} onlineDiv={onlineDiv} lastMessage={last_message} />
                        </div>
                    );
                })}
            </div>
        </div>
    );
}