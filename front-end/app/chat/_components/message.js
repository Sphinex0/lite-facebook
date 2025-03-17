
import styles from "../styles.module.css";
export default function Message(msg, index) {
    return (
        <div key={`msg-${index}`} className={styles.message}>
            <div className={styles.messageHeader}>
                <span className={styles.userName}>
                    {msg.sender_id}
                </span>
            </div>
            <div className={styles.messageContent}>
                {msg.content}
            </div>
        </div>
    )
}