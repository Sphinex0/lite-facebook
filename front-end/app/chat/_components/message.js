
import styles from "../styles.module.css";
export default function Message({ msg }) {
    const { message, user_info } = msg

    return (
        <>
            {
                message && <div className={styles.message}>
                    <div className={styles.messageHeader}>
                        <span className={styles.userName}>
                            {user_info.first_name} {user_info.last_name}
                        </span>
                    </div>
                    <div className={styles.messageContent}>
                        {message.content}
                    </div>
                </div>
            }
        </>
    )
}