import { ThumbUp, ThumbDown } from "@mui/icons-material";
import styles from './post.module.css';

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
      <div className="post-details">
         <h2>{postInfo.article.content}</h2>
         {postInfo.article.image && <img src={postInfo.article.image} alt="Post image" />}
      
         {/* Like and Dislike buttons */}
         <div className="action-button">
           <div className="action-buttons">
             <span>
               <ThumbUp
                 onClick={() => likePost(1, postInfo.article.id)}
                 className={`${likeState === 1 ? styles.blue : ""} ${styles.ArticleActionBtn}`}
               />
               <span className={styles.footerText}>{likes}</span>
               <ThumbDown
                 onClick={() => likePost(-1, postInfo.article.id)}
                 className={`${likeState === -1 ? styles.red : ""} ${styles.ArticleActionBtn}`}
               />
               <span className={styles.footerText}>{disLikes}</span>
  
               <span className={styles.footerText}>{commentsCount}</span>
             </span>
           </div>
         </div>
      </div>
      <form
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
      </form>
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