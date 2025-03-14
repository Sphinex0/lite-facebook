'use client'
import { Comment, Preview, PrivacyTip, Public, ThumbDown, ThumbUp } from "@mui/icons-material"
import { use, useEffect, useState } from "react"
import styles from './post.module.css'
import PostViewer from "./postViewer"
import { likeArticle, timeAgo } from "@/app/helpers"

export default function Post({ postInfo }) {
    const [likes, setLikes] = useState(postInfo.likes || 0); // Fallback to 0 if undefined
    const [disLikes, setDislikes] = useState(postInfo.disLikes || 0);
    const [commentsCount, setCommentCount] = useState(postInfo.comments_count || 0);
    const [likeState, setLikeState] = useState(postInfo.like || 0);
    const [postViewDisplay, setPostViewDisplay] = useState(false);
  

    const likePost = (like, article_id)=>{
        likeArticle(like, article_id, setLikes,setDislikes, likeState, setLikeState)
    }


    return (
        <div className="feed">
            <div className="head">
                <div className="user">
                    <div className="profile-photo">
                        <img src="./images/profile-13.jpg" />
                    </div>
                    <div className="ingo">
                        <h3>{postInfo.user_info.first_name} {postInfo.user_info.last_name}</h3>
                        <small>{timeAgo(postInfo.article.created_at)}</small>
                    </div>
                </div>
            </div>
            
            <Public/> {postInfo.article.privacy}
            <div  className={styles.content}>{postInfo.article.content}</div>


            {postInfo.article.image && <img src="./images/feed-1.jpg" />}

            <div className="action-button">
                <div className="action-buttons">
                    <span>
                        <ThumbUp onClick={() => { likePost(1, postInfo.article.id) }} className={`${likeState == 1 ? styles.blue : ""} ${styles.ArticleActionBtn}`} />
                        <span className={styles.footerText}>{likes}</span>

                        <ThumbDown onClick={() => { likePost(-1, postInfo.article.id) }} className={`${likeState == -1 ? styles.red : ""} ${styles.ArticleActionBtn}`} />
                        <span className={styles.footerText}>{disLikes}</span>
                    </span>

                    <span>
                        <Comment className={styles.ArticleActionBtn} onClick={() => setPostViewDisplay(true)} />
                        <span className={styles.footerText}>{commentsCount}</span>
                        {postViewDisplay && (<PostViewer
                            postInfo={postInfo}
                            likes={likes}
                            disLikes={disLikes}
                            likeState={likeState}
                            likePost={likePost}
                            commentsCount={commentsCount}
                            setPostViewDisplay={setPostViewDisplay}
                        />
                        )}
                    </span>
                </div>
            </div>
        </div>
    )
}