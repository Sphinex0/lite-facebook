import { ThumbUp, ThumbDown, Comment } from "@mui/icons-material";
import styles from './post.module.css';
import { timeAgo } from "@/app/helpers";
import CreateComment from "./createComment";

export default function PostViewer({ postInfo, likes, disLikes, likeState, likePost, commentsCount, setPostViewDisplay }) {
  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setPostViewDisplay(false)
    }
  }
  return (
    // <div className="post-viewer">
    //   {/* Existing content, e.g., post details or comments */}
    //   <h2>{postInfo.article.content}</h2>

    //   {/* Like and Dislike buttons */}
    //   <div className="action-button">
    //     <div className="action-buttons">
    //       <span>
    //         <ThumbUp
    //           onClick={() => likePost(1, postInfo.article.id)}
    //           className={`${likeState === 1 ? styles.blue : ""} ${styles.ArticleActionBtn}`}
    //         />
    //         <span className={styles.footerText}>{likes}</span>
    //         <ThumbDown
    //           onClick={() => likePost(-1, postInfo.article.id)}
    //           className={`${likeState === -1 ? styles.red : ""} ${styles.ArticleActionBtn}`}
    //         />
    //         <span className={styles.footerText}>{disLikes}</span>

    //         <span className={styles.footerText}>{commentsCount}</span>
    //       </span>
    //     </div>
    //   </div>
    // </div>
    <div className="customize-theme" onClick={hide}>
      <div className="card">
        <h2>Post</h2>
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
          <p>
            {postInfo.article.content}
          </p>
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
        {/* <form
          className={styles.form}
          onSubmit={(e) => {
            //   e.preventDefault();
            //   addComment(commentContent);
            //   setCommentContent('');
          }}
        >
          <textarea
            className={styles.textInput}
            //value={commentContent}
            //onChange={(e) => setCommentContent(e.target.value)}
            placeholder="Write a comment..."
          ></textarea>
          <button type="submit" className="btn btn-primary">Comment</button>
        </form> */}
        <CreateComment />
        <div className="comments">
          {/*comments.map((comment) => (
          <div key={comment.id} className="comment">
            <p>{comment.author}</p>
            <div style={{ whiteSpace: 'pre-wrap' }}>{comment.content}</div>
          </div>
        ))*/}
        </div>
      </div>
    </div>
  );
}