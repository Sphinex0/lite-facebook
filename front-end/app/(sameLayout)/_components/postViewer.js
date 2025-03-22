import { ThumbUp, ThumbDown, OpenInNew } from "@mui/icons-material";
import styles from './post.module.css';
import { FetchApi, timeAgo, useOnVisible } from "@/app/helpers";
import CreateComment from "./createComment";
import { useEffect, useRef, useState } from "react";
import Comment from "./comment";
import UserInfo from "./userInfo";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function PostViewer({ postInfo, likes, disLikes, likeState, likePost, commentsCount, setCommentCount, setPostViewDisplay }) {
  const [comments, setComments] = useState([])
  const lastElementRef = useRef(null)
  const before = useRef(Math.floor(Date.now()))
  const redirect = useRouter()

  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setPostViewDisplay(false)
    }
  }


  const fetchComments = async (signal) => {
    try {
      const response = await FetchApi("/api/comments", redirect, {
        method: "POST",
        body: JSON.stringify({ before: before.current, parent: postInfo.article.id }),
        signal
      })

      console.log("status:", response.status)
      if (response.ok) {
        const commentsData = await response.json()
        console.log(commentsData)
        if (commentsData) {


          setComments((prv) => [...prv, ...commentsData])


          before.current = commentsData[commentsData.length - 1].article.created_at
          console.log("last created at", commentsData[commentsData.length - 1].article.created_at)
        }
      }

    } catch (error) {
      console.log(error)
    }

  }
  useEffect(() => {
    const controller = new AbortController()
    fetchComments(controller.signal)
    return () => controller.abort()
  }, [])
  useOnVisible(lastElementRef, fetchComments)


  return (
    <div className="customize-theme" onClick={hide}>
      <div className="card">
        <h2>Post</h2>
        <div className="feed">
          {postInfo.group_name ? <div> <strong> Group </strong>: {<Link href={`/groups/${postInfo.article.group_id}`}>{postInfo.group_name}</Link>}</div> : ""}
          <div className="head">
            <UserInfo userInfo={postInfo.user_info} articleInfo={postInfo.article} />
          </div>
          <div className={`${styles.content} ${styles.PreviewContent}`}>{postInfo.article.content}</div>

          {postInfo.article.image && <div className={styles.imageHolder}><img src={`/posts/${postInfo.article.image}`} /> <a href={`/posts/${postInfo.article.image}`} target="_blank" className={styles.OpenInNew}><OpenInNew /></a> </div>}

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
        <CreateComment setComments={setComments} setCommentCount={setCommentCount} parent={postInfo.article.id} />
        <div className="comments">
          {comments.length === 0 ? <h5> no comments yet</h5> : comments.map((comment) => {
            return <Comment key={comment.article.id} commentInfo={comment} reference={lastElementRef} />
          }

          )}

        </div>
      </div>
    </div>
  );
}