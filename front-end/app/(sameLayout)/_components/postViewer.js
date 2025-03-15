import { ThumbUp, ThumbDown } from "@mui/icons-material";
import styles from './post.module.css';
import { timeAgo } from "@/app/helpers";
import CreateComment from "./createComment";
import { useEffect, useRef, useState } from "react";
import Comment from "./comment";
import UserInfo from "./userInfo";

export default function PostViewer({ postInfo, likes, disLikes, likeState, likePost, commentsCount, setPostViewDisplay }) {
  const [comments, setComments] = useState([])
  const before = useRef(Math.floor(Date.now() / 1000))

  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setPostViewDisplay(false)
    }
  }


  const fetchComments = async (first = false) => {
    try {
      const response = await fetch("http://localhost:8080/api/comments", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({ before: before.current, parent: postInfo.article.id })
      })

      console.log("status:", response.status)
      if (response.ok) {
        const commentsData = await response.json()
        console.log(commentsData)
        if (commentsData) {

          if (first) {
            setComments(commentsData)
          } else {
            setComments((prv) => [...prv, ...commentsData])
          }

          before.current = commentsData[commentsData.length - 1].article.created_at
          console.log("last created at", commentsData[commentsData.length - 1].article.created_at)
        }
      }

    } catch (error) {
      console.log(error)
    }

  }
  useEffect(() => {
    console.log("fetch coments")
    fetchComments(true)
  }, [])
  return (
    <div className="customize-theme" onClick={hide}>
      <div className="card">
        <h2>Post</h2>
        <div className="feed">
          <div className="head">
            <UserInfo userInfo={postInfo.user_info} articleInfo={postInfo.article}/>
          </div>
          <div className={`${styles.content} ${styles.PreviewContent}`}>{postInfo.article.content}</div>

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
                <span className={styles.footerText}>{commentsCount} Comments</span>
              </span>
            </div>
          </div>
        </div>

        <p style={{ textAlign: "left" }}>Comments :</p>
        <CreateComment setComments={setComments} parent={postInfo.article.id} />
        <div className="comments">
          {comments.length === 0 ? <h5> no comments yet</h5> : comments.map((comment) => {
            // <div key={comment.article.id} className="comment">
            //   <p>{comment}</p>
            //   <div style={{ whiteSpace: 'pre-wrap' }}>{comment}</div>
            // </div>
            return <Comment key={comment.article.id} commentInfo={comment}/>
          }

          )}
        </div>
      </div>
    </div>
  );
}