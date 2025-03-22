"use client";
import { useEffect, useState, useRef, use } from "react";
import styles from "./styles.module.css";
import Message from "@/app/(sameLayout)/chat/_components/message";
import { AddPhotoAlternate, Cancel, EmojiEmotions, PeopleAlt, Send } from "@mui/icons-material";
import { emojis } from "./_components/emojis";
import UserInfo from "../_components/userInfo";
import { FetchApi, opThrottle } from "@/app/helpers";
import { useWorker } from "@/app/_Context/WorkerContext";
import { useRouter } from "next/navigation";

export default function Chat() {

    const { portRef, clientWorker, conversations, setConversations, selectedConversationRef, messages, setMessages } = useWorker();
    const conversationsRef = useRef(conversations);
    const [message, setMessage] = useState({ content: "", reply: null });
    const [selectedConversation, setSelectedConversation] = useState(null);
    const [img, setImage] = useState(null);
    const [emoji, setEmoji] = useState(false);
    const chatEndRef = useRef(null);
    const beforeRef = useRef(Math.floor(new Date().getTime()));
    const combinadeRef = useRef(null);
    const [replyingTo, setReplyingTo] = useState(null);
    const inputRef = useRef(null);
    const chatBodyRef = useRef(null);
    const oldScrollHeightRef = useRef(null); 
    const justLoadedMoreRef = useRef(false); 
    const redirect = useRouter()


    useEffect(() => {
        conversationsRef.current = conversations;
    }, [conversations]);

    useEffect(() => {
        const port = portRef.current;
        if (!port) return;
        port.postMessage({
            kind: "send",
            payload: {
                type: "conversations"
            }
        });
    }, [])

    const fetchMessages = async (signal) => {
        if (!selectedConversation) return;
        const res = await FetchApi(`/api/messageshistories`,redirect, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ conversation_id: selectedConversation.id, before: beforeRef.current }),
            signal
        });



        if (res.ok) {
            const data = await res.json() || [];
            if (data && data.length > 0) {
                const revercedData = data.reverse();
                setMessages((prev) => [...revercedData, ...prev]);
                beforeRef.current = data[0].message.created_at

            }
        } else {
            console.error("Error fetching messages");
        }
    };

    const handleScroll = (event) => {
        if (event.target.scrollTop === 0 && messages.length > 0) {
            oldScrollHeightRef.current = chatBodyRef?.current?.scrollHeight;
            justLoadedMoreRef.current = true;
            const controller = new AbortController();
            opThrottle(fetchMessages, 1000)(controller.signal);
            return () => {
                controller.abort();
            };
        }
    };


    useEffect(() => {
        selectedConversationRef.current = selectedConversation;
        const controller = new AbortController();
        fetchMessages(controller.signal);

        return () => {
            controller.abort();
        }
    }, [selectedConversation]);

    useEffect(() => {
        if (justLoadedMoreRef.current) {
            const chatBody = chatBodyRef.current;
            if (chatBody) {
                const newScrollHeight = chatBody.scrollHeight;
                const deltaH = newScrollHeight - oldScrollHeightRef.current;
                chatBody.scrollTop = deltaH;
            }
            justLoadedMoreRef.current = false;
        } else {
            chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
        }
    }, [messages]);


    const handleReply = (message) => {
        if (replyingTo?.id === message.message.id) {
            // Clear reply if clicking the same message
            setReplyingTo(null);
            setMessage(prev => ({ ...prev, reply: null }));
        } else {
            // Set new reply
            setReplyingTo({
                id: message.message.id,
                content: message.message.content,
                sender: message.user_info?.first_name || "Unknown"
            });
            setMessage(prev => ({ ...prev, reply: message.message.id }));
        }
        inputRef.current?.focus();
    };

    const cancelReply = () => {
        setReplyingTo(null);
        setMessage(prev => ({ ...prev, reply: null }));
    };

    const handleSendMessage = (event) => {
        if (event.key !== "Enter" || !message.content.trim()) return;
        const port = portRef.current;
        if (!port || !selectedConversation) return;

        port.postMessage({
            kind: "send",
            payload: {
                type: "new_message",
                message: {
                    conversation_id: selectedConversation.id,
                    content: message.content,
                    reply: message.reply
                },
            },
        });

        setMessage({ content: "", reply: null });
        setReplyingTo(null);
        setEmoji(false);
    };



    const SendFile = () => {
        const port = portRef.current;
        if (combinadeRef.current) {
            port.postMessage({
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
        const port = portRef.current;
        if (selectedConversation?.id !== conversation.id) {
            const type = conversation.type == "private" ? "read_messages_private" : "read_messages_group";
            port.postMessage({
                kind: "send",
                payload: {
                    type,
                    message: {
                        conversation_id: conversation.id,
                    },
                },
            });
            setConversations((prev) => {
                return prev.map((c) => {
                    if (c.conversation.id === conversation.id) {
                        return {
                            ...c,
                            seen: 0
                        };
                    }
                    return c;
                });
            });
            setMessages([]);
            cancelReply()
            beforeRef.current = Math.floor(new Date().getTime());
            setSelectedConversation(conversation);
        }
    };

    const selectedConversationInfo = conversations?.find(
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
                    <div>
                        <label htmlFor="Hide" className={styles.ConversationsBtn}>
                            <PeopleAlt />
                        </label>
                    </div>
                    <h4>{displayTitle}</h4>
                </div>


                <div className={styles.chatBody} onScroll={handleScroll} ref={chatBodyRef}>
                    {selectedConversation ? (
                        messages.length > 0 ? (
                            <>
                                {messages.map((msg, index) => (
                                    <Message
                                        msg={msg}
                                        key={msg.message.id}
                                        onClick={() => handleReply(msg)}
                                        isSelected={replyingTo?.id === msg.message.id}
                                    />
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

                {emoji && selectedConversation && (
                    <div className={styles.Emojis}>
                        {emojis.map((emo) => (
                            <div
                                className={styles.Emoji}
                                key={emo}
                                onClick={() => setMessage((prev) => {
                                    return {
                                        ...prev,
                                        content: prev.content + emo
                                    }
                                })}
                            >
                                {emo}
                            </div>
                        ))}
                    </div>
                )}

                {/* Reply Preview Bar */}
                {replyingTo && (
                    <div className={styles.replyPreviewBar}>
                        <div className={styles.replyPreviewContent}>
                            <div className={styles.replyPreviewTitle}>
                                Replying to {replyingTo.sender}
                            </div>
                            <div className={styles.replyPreviewText}>
                                {replyingTo.content}
                            </div>
                        </div>
                        <Cancel
                            className={styles.replyCancel}
                            onClick={cancelReply}
                            fontSize="small"
                        />
                    </div>
                )}

                <div className={styles.groupInputs}>

                    <input
                        ref={inputRef}
                        className={styles.chatInput}
                        value={message.content}
                        onChange={(e) => setMessage(prev => ({ ...prev, content: e.target.value }))}
                        onKeyDown={handleSendMessage}
                        placeholder={replyingTo ? "Type your reply..." : "Type your message..."}
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

            <input type="checkbox" id="Hide" className={styles.Hide} />
            <div className={styles.conversationsList}>
                
                {conversations?.map((conversationInfo) => {
                    const { conversation, user_info, group, last_message, seen } = conversationInfo;
                    const onlineDiv = true
                    return (
                        <div
                            key={`conv-${conversation.id}`}
                            className={`${styles.conversationItem} ${selectedConversation?.id === conversation.id ? styles.active : ""
                                }`}
                            onClick={() => handleSetSelectedConversation(conversation)}
                        >
                            <div>
                                <UserInfo userInfo={user_info} group={group} onlineDiv={onlineDiv} lastMessage={last_message} />
                            </div>
                            <div >
                                <div className={styles.seen}>{seen}</div>
                            </div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}
