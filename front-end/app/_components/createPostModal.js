'use client'
import { AddPhotoAlternate, Public, SentimentDissatisfiedSharp } from '@mui/icons-material'
import React, { useEffect, useState } from 'react'
import styles from './createPostModal.module.css'

const CreatePostModal = ({ setModalDisplay }) => {
  const [content, setContent] = useState("")
  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setModalDisplay(false)
    }
  }

  const addPost = (e) => {
    e.preventDefault()
    const formData = new FormData(e.target)
    console.log(formData)
    setContent("")
  }

  useEffect(() => {
    document.querySelector(`.${styles.textInput}`).textContent = content
  }, [content])

  return (
    <div className="customize-theme" onClick={hide}>
      <div className="card">
        <h2>Create post</h2>
        <form action="" className={styles.form} onSubmit={addPost}>
          <select name='privacy' className={styles.selectPrivacy}>
            <option value="public"> Public</option>
            <option value="almost_private">Almost Private</option>
            <option value="private">Private</option>
          </select>

          {/* <div contentEditable className={styles.textInput} 
          onBlur={()=>{content == "" ? setContent("What's on your mind, Diana ?"):""}} 
          onFocus={()=>content == "What's on your mind, Diana ?" ? setContent(""):""} 
          onInput={(e)=>{setContent(e.target.textContent)}}>
              
          </div> */}
          <textarea className={styles.textInput}
            placeholder={"What's on your mind, Diana ?"}
            value={content}
            onInput={(e) => { setContent(e.target.textContent) }}>
          </textarea>

          <input type="file" id='postImage' className={styles.fileInput} />

          <div className={styles.footer}>
            <label htmlFor="postImage">
              <AddPhotoAlternate />
            </label>

            <input type="submit" value="Post" className="btn btn-primary" />
          </div>
        </form>
      </div>
    </div>
  )
}

export default CreatePostModal
