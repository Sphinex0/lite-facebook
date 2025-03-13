'use client'
import { AddPhotoAlternate, Public, SentimentDissatisfiedSharp } from '@mui/icons-material'
import React, { useEffect, useState } from 'react'
import styles from './createPostModal.module.css'

const CreatePostModal = ({ setModalDisplay , setPosts}) => {
  const [content, setContent] = useState("")
  const [imagePreview, setImagePreview] = useState("")

  const hide = (e) => {
    if (e.target.classList.contains('customize-theme')) {
      setModalDisplay(false)
    }
  }

  const addPost = async(e) => {
    e.preventDefault()
    try {
      const formData = new FormData(e.target)
      console.log(formData)
      const response = await fetch("http://localhost:8080/api/articles/store", {
          method: "POST",
          credentials: "include",
          body: formData
      })

      console.log("status:", response.status)
      if (response.ok) {
          const article = await response.json()
          setPosts((prv)=>[{article,user_info:{}},...prv])
          setModalDisplay(false) 
          setContent("")
      }

  } catch (error) {
      console.log(error)
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
          <textarea name='content' className={styles.textInput}
            placeholder={"What's on your mind, Diana ?"}>
          </textarea>

          {imagePreview && <img src={imagePreview} />}
          <input 
          type="file" 
          id='postImage'
          onChange={(e)=>{
            if (e.target.files[0]){
              const file = e.target.files[0]
              const reader = new FileReader()
              reader.onloadend=()=>{
                setImagePreview(reader.result)
              }
              reader.readAsDataURL(file)
            }
          }}
          className={styles.fileInput} />

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
