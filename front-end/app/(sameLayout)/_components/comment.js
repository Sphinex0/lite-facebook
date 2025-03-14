import React from 'react'
import styles from "./comment.module.css"
import { ThumbDown, ThumbUp } from '@mui/icons-material';
import { likeArticle } from '@/app/helpers';
const Comment = ({ commentInfo }) => {
  const [likes, setLikes] = useState(postInfo.likes || 0); // Fallback to 0 if undefined
  const [disLikes, setDislikes] = useState(postInfo.disLikes || 0);
  const [likeState, setLikeState] = useState(postInfo.like || 0);


  return (
    <div key={commentInfo.id} className="comment">
      <p>{commentInfo.author}</p>
      <div className={styles.content}>{commentInfo.content}</div>

      {commentInfo.image && <img src="./images/feed-1.jpg" />}


      <div className="action-button">
        <div className="action-buttons">
          <span>
            <ThumbUp onClick={() => { likeArticle(1, commentInfo.id) }} className={`${likeState == 1 ? styles.blue : ""} ${styles.ArticleActionBtn}`} />
            <span className={styles.footerText}>{likes}</span>

            <ThumbDown onClick={() => { likeArticle(-1, commentInfo.id) }} className={`${likeState == -1 ? styles.red : ""} ${styles.ArticleActionBtn}`} />
            <span className={styles.footerText}>{disLikes}</span>
          </span>


        </div>
      </div>
    </div>
  )
}

export default Comment
