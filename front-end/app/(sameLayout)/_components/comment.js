import React, { useState } from 'react'
import styles from "./comment.module.css"
import Pstyles from "./post.module.css"
import { ThumbDown, ThumbUp } from '@mui/icons-material';
import { likeArticle } from '@/app/helpers';
import UserInfo from './userInfo';
import { useRouter } from 'next/navigation';
const Comment = ({ commentInfo , reference}) => {
  const [likes, setLikes] = useState(commentInfo.likes || 0); // Fallback to 0 if undefined
  const [disLikes, setDislikes] = useState(commentInfo.disLikes || 0);
  const [likeState, setLikeState] = useState(commentInfo.like || 0);
  const redirect = useRouter()
  console.log(commentInfo)
  
  return (
    <div className={styles.comment} ref={reference}>
      <UserInfo userInfo={commentInfo.user_info} articleInfo={commentInfo.article}/>
      <div className={styles.content}>{commentInfo.article.content}</div>

      {commentInfo.article.image && <img src={`/posts/${commentInfo.article.image}`} />}


      <div className="action-button">
        <div className="action-buttons">
          <span>
            <ThumbUp onClick={() => { likeArticle(1, commentInfo.article.id,setLikes,setDislikes,likeState,setLikeState ,redirect) }} className={`${likeState == 1 ? Pstyles.blue : ""} ${Pstyles.ArticleActionBtn}`} />
            <span className={styles.footerText}>{likes}</span>

            <ThumbDown onClick={() => { likeArticle(-1, commentInfo.article.id,setLikes,setDislikes,likeState,setLikeState,redirect ) }} className={`${likeState == -1 ? Pstyles.red : ""} ${Pstyles.ArticleActionBtn}`} />
            <span className={styles.footerText}>{disLikes}</span>
          </span>


        </div>
      </div>
    </div>
  )
}

export default Comment
