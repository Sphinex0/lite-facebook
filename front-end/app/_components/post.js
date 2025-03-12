import { Comment, ThumbDown, ThumbUp } from "@mui/icons-material"
import { use, useState } from "react"
import styles from './post.module.css'

export default function Post({ postInfo }) {
    const [likes, setLikes] = useState(postInfo.likes)
    const [disLikes, setDislikes] = useState(postInfo.disLikes)
    const [commentsCount, setCommentCount] = useState(postInfo.comments_count)
    const [likeState, setLikeState] = useState(postInfo.like) //0 1 -1
    console.log(postInfo)
    const likePost = async (like, article_id) => {
        try {
            const response = await fetch("http://localhost:8080/api/reactions/store", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ like, article_id })
            })

            console.log("status:", response.status)
            if (response.ok) {
                console.log("liked or dislike")
                if (like == 1 && likeState == 1) {
                    setLikes(likes - 1)
                    setLikeState(0)
                } else if (like == -1 && likeState == -1) {
                    setDislikes(disLikes - 1)
                    setLikeState(0)
                } else if (like == 1 && likeState == -1) {
                    setDislikes(disLikes - 1)
                    setLikes(likes + 1)
                    setLikeState(1)
                } else if (like == -1 && likeState == 1) {
                    setDislikes(disLikes + 1)
                    setLikes(likes - 1)
                    setLikeState(-1)
                } else if (like == 1 && likeState == 0) {
                    setLikes(likes + 1)
                    setLikeState(1)
                } else if (like == -1 && likeState == 0) {
                    setDislikes(disLikes + 1)
                    setLikeState(-1)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }
    const pluralize = (number,unit)=>{
        if (number === 1) {
            return `${number} ${unit} ago`;
        } else {
            return `${number} ${unit}s ago`;
        }
    }
    const timeAgo = (unixTimestamp)=>{
        // Convert Unix timestamp (in seconds) to a Date object (requires milliseconds)
    const pastDate = new Date(unixTimestamp * 1000);
    const currentDate = new Date();

    // Calculate the difference in milliseconds
    const diffMs = currentDate - pastDate;

    // Handle future timestamps
    if (diffMs < 0) {
        return "in the future";
    }

    // Convert difference to seconds
    const diffSeconds = Math.floor(diffMs / 1000);

    // Define time thresholds and calculate appropriate unit
    if (diffSeconds < 60) {
        // Less than 1 minute
        if (diffSeconds < 10) {
            return "just now"; // Very recent, less than 10 seconds
        } else {
            return pluralize(diffSeconds, "second");
        }
    } else if (diffSeconds < 3600) {
        // Less than 1 hour (3600 seconds)
        const minutes = Math.floor(diffSeconds / 60);
        return pluralize(minutes, "minute");
    } else if (diffSeconds < 86400) {
        // Less than 1 day (86400 seconds)
        const hours = Math.floor(diffSeconds / 3600);
        return pluralize(hours, "hour");
    } else if (diffSeconds < 2592000) {
        // Less than 30 days (2592000 seconds)
        const days = Math.floor(diffSeconds / 86400);
        return pluralize(days, "day");
    } else if (diffSeconds < 31536000) {
        // Less than 1 year (31536000 seconds, approximating 365 days)
        const months = Math.floor(diffSeconds / 2592000); // Approximate month as 30 days
        return pluralize(months, "month");
    } else {
        // 1 year or more
        const years = Math.floor(diffSeconds / 31536000); // Approximate year as 365 days
        return pluralize(years, "year");
    }
    }


    /**
     *             <hr/>
          <h1>content : {postInfo.article.content}</h1>
          <h3><span onClick={()=>{likePost(1, postInfo.article.id)}}>likes</span>: {likes} || <span onClick={()=>{likePost(-1, postInfo.article.id)}}>dislikes</span>: {disLikes}</h3>
          <hr/>
     */
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
            <div className="photo">
                {postInfo.article.content}
                {/*<img src="./images/feed-1.jpg" />*/}
            </div>
            <div className="action-button">
                <div className="action-buttons">
                    <span>
                        <ThumbUp onClick={()=>{likePost(1, postInfo.article.id)}} className={`${likeState == 1 ? styles.blue: ""} ${styles.ArticleActionBtn}`}/>
                        <span className={styles.footerText}>{likes}</span>

                        <ThumbDown onClick={()=>{likePost(-1, postInfo.article.id)}} className={`${likeState == -1 ? styles.red: ""} ${styles.ArticleActionBtn}`}/> 
                        <span className={styles.footerText}>{disLikes}</span>
                    </span>
                    <span>
                        <Comment className={styles.ArticleActionBtn}/> 
                        <span className={styles.footerText}>{commentsCount}</span>
                    </span>
                </div>
            </div>
        </div>
    )
}