
import Image from "next/image";
import styles from "../styles.module.css";
export default function Message({ msg }) {
    const { message, user_info } = msg
    // message.content = `\uD83D\uDE00`
    return (
        <>
            {
                message && <div className={styles.message}>
                    <div className={styles.messageHeader}>
                        <span className={styles.userName}>
                            {user_info.first_name} {user_info.last_name}
                        </span>
                    </div>
                    {
                        message.content && <div className={styles.messageContent}>
                            {message.content}
                        </div>
                    }

                    {
                        message.image && <div >
                            <img className={styles.messageImage} src={`/images/${message.image}`} />
                        </div>
                    }
                </div>
            }
        </>
    )
}