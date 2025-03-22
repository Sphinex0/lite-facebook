import { Comment, Preview, PrivacyTip, Public, ThumbDown, ThumbUp } from "@mui/icons-material"
import { use, useEffect, useState } from "react"
import styles from './post.module.css'
import PostViewer from "./postViewer"
import { likeArticle, timeAgo } from "@/app/helpers"
import UserInfo from "./userInfo"
import Link from "next/link"
import { useRouter } from "next/navigation"

export default function Post({ postInfo , reference}) {
    const [likes, setLikes] = useState(postInfo.likes || 0); // Fallback to 0 if undefined
    const [disLikes, setDislikes] = useState(postInfo.disLikes || 0);
    const [commentsCount, setCommentCount] = useState(postInfo.comments_count || 0);
    const [likeState, setLikeState] = useState(postInfo.like || 0);
    const [postViewDisplay, setPostViewDisplay] = useState(false);
    const redirect = useRouter()
  

    const likePost = (like, article_id)=>{
        likeArticle(like, article_id, setLikes,setDislikes, likeState, setLikeState,redirect)
    }


    return (
        <div className="feed" ref={reference}>
            
            {postInfo.group_name ?<div> <strong> Group </strong>: {<Link href={`/groups/${postInfo.article.group_id}`}>{postInfo.group_name}</Link>}</div> : ""}
            <div className="head">
                <UserInfo userInfo={postInfo.user_info} articleInfo={postInfo.article}/>
            </div>
            
            
            <div className={styles.content}>{postInfo.article.content}</div>


            {postInfo.article.image && <img src={`/posts/${postInfo.article.image}`} />}

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
                            setCommentCount={setCommentCount}
                            setPostViewDisplay={setPostViewDisplay}
                        />
                        )}
                    </span>
                </div>
            </div>
        </div>
    )
}