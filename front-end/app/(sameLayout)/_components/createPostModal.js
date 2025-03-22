'use client'
import { AddPhotoAlternate, Public, SentimentDissatisfiedSharp } from '@mui/icons-material'
import React, { useEffect, useRef, useState } from 'react'
import styles from './createPostModal.module.css'
import SelectFollower from './selectFollower'
import { addArticle } from '@/app/helpers'
import { useRouter } from 'next/navigation'

const CreatePostModal = ({ setModalDisplay, setPosts , group}) => {
  const [content, setContent] = useState("")
  const [imagePreview, setImagePreview] = useState("")
  const [privacy, setPrivacy] = useState("")
  const redirect = useRouter()

  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setModalDisplay(false)
    }
  }

  const addPost = async (e) => {
    const added = await addArticle(e, setPosts, {},redirect)
    if (added) {
      setModalDisplay(false)
      setContent("")
      setImagePreview("")
    }


  }

  useEffect(() => {
    document.querySelector(`.${styles.textInput}`).textContent = content
  }, [content])

  return (
    <div className="customize-theme" onClick={hide}>
      <div className="card">
        <h2>Create post</h2>
        <form action="" className={styles.form} onSubmit={addPost}>
          <select name='privacy' className={styles.selectPrivacy} onChange={(e) => setPrivacy(e.target.value)}>
            <option value="public"> Public</option>
            <option value="almost_private">Almost Private</option>
            <option value="private">Private</option>
          </select>
          {privacy === "private" && <SelectFollower />}
          <textarea name='content' className={styles.textInput}
            placeholder={"What's on your mind, Diana ?"}>
          </textarea>

          <input
            type="file"
            id='postImage'
            name='image'
            onChange={(e) => {
              if (e.target.files[0]) {
                const file = e.target.files[0]
                const reader = new FileReader()
                reader.onloadend = () => {
                  setImagePreview(reader.result)
                }
                reader.readAsDataURL(file)
              } else {
                setImagePreview("")
              }
            }}
            className={styles.fileInput} />

          <div className={styles.footer}>
            <label htmlFor="postImage" className={styles.addFile}>
            {imagePreview ? <img src={imagePreview} className="imagePreview" /> : <AddPhotoAlternate />}
            </label>

            <input type="submit" value="Post" className="btn btn-primary" />
          </div>
        </form>
      </div>
    </div>
  )
}

export default CreatePostModal