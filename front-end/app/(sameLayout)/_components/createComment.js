import React, { useState } from 'react'
import styles from "./comment.module.css"
// import styles from "./createPostModal.module.css"
import { AddPhotoAlternate } from '@mui/icons-material'

const CreateComment = () => {
    const [imagePreview, setImagePreview] = useState("")

    return (
        <>
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
            <input
                type="file"
                id='postImage'
                onChange={(e) => {
                    if (e.target.files[0]) {
                        const file = e.target.files[0]
                        const reader = new FileReader()
                        reader.onloadend = () => {
                            setImagePreview(reader.result)
                        }
                        reader.readAsDataURL(file)
                    }
                }}
                className={styles.fileInput} />
            <label htmlFor="postImage" className={styles.addFile}>
                <AddPhotoAlternate />
            </label>
            <button type="submit" className="btn btn-primary">Comment</button>
        </form>
        {imagePreview && <img src={imagePreview} className={styles.imagePreview}/>}
        </>
    )
}

export default CreateComment
