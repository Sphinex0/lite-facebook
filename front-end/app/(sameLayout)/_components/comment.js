import React from 'react'
import styles from "./comment.module.css"
const comment = ({commentInfo}) => {
    return (
        <div key={commentInfo.id} className="comment">
            <p>{commentInfo.author}</p>
            <div className={styles.content}>{commentInfo.content}</div>

            {postInfo.article.image && <img src="./images/feed-1.jpg" />}


            <div className="action-button">
            <div className="action-buttons">
              <span>
                <ThumbUp onClick={() => { likePost(1, postInfo.article.id) }} className={`${likeState == 1 ? styles.blue : ""} ${styles.ArticleActionBtn}`} />
                <span className={styles.footerText}>{likes}</span>

                <ThumbDown onClick={() => { likePost(-1, postInfo.article.id) }} className={`${likeState == -1 ? styles.red : ""} ${styles.ArticleActionBtn}`} />
                <span className={styles.footerText}>{disLikes}</span>
              </span>


            </div>
          </div>
        </div>
    )
}

export default comment
