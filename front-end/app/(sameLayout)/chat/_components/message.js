
"use client";
import styles from "../styles.module.css";

export default function Message({ msg, onClick, isSelected }) {
    return (
        <div
            className={`${styles.messageContainer} ${isSelected ? styles.selectedMessage : ''}`}
            onClick={onClick}
        >
            {/* Reply Reference */}
            {
                msg.reply_content ? (
                    <div className={styles.replyReference}>
                        <div className={styles.replyPreviewText}>
                            Replying to: {msg.reply_content}
                        </div>
                    </div>
                ) : ""
            }

            <div className={styles.messageMeta}>
                <span className={styles.name}>
                    {msg.user_info.first_name}
                </span>
            </div>

            {/* Message Content */}
            <div className={styles.messageContent}>
                {
                    msg.message.content || <img className={styles.messageImage} src={`/images/${msg.message.image}`} />
                }
            </div>

            {/* Metadata */}
            <div className={styles.messageMeta}>
                <span className={styles.timestamp}>
                    {console.log(msg.message.created_at)}
                    {new Date(msg.message.created_at).toLocaleTimeString()}
                </span>
            </div>
        </div>
    );
}