
// import Image from "next/image";
// import styles from "../styles.module.css";
// import ReplyIcon from '@mui/icons-material/Reply';
// export default function Message({ msg , onClick }) {
//     const { message, user_info } = msg
//     // message.content = `\uD83D\uDE00`
//     return (
//         <div onClick={() => onClick(msg)} >
//             {
//                 message && <div className={styles.message}>
//                     <div className={styles.messageHeader}>
//                         <span className={styles.userName}>
//                             {user_info.first_name} {user_info.last_name}
//                         </span>
//                         <small className={styles.ReplyIcon}> <ReplyIcon  /> </small>
//                     </div>
//                     {
//                         message.content && <div className={styles.messageContent}>
//                             {message.content}
//                         </div>
//                     }

//                     {
//                         message.image && <div >
//                             <img className={styles.messageImage} src={`/images/${message.image}`} />
//                         </div>
//                     }
//                 </div>
//             }
//         </div>
//     )
// }


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