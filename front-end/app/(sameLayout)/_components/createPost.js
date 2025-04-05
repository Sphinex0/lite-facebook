'use client'
import { useWorker } from '@/app/_Context/WorkerContext'
import React, { useEffect } from 'react'

const CreatePost = ({ setModalDisplay }) => {
  const show = () => {
    setModalDisplay(true)
  }
  const {userRef} = useWorker()

  return (
    <div className="create-post" onClick={show}>
      <div className="profile-photo">
        {userRef.current && <img src={`http://localhost:8080/public/pics/${userRef.current.image}`} />}
      </div>
      <div id="create-post" >What's on your mind, Diana ?</div>
    </div>
  )
}

export default CreatePost
