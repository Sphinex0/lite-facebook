import React, { useEffect } from 'react'

const CreatePost = ({ setModalDisplay }) => {
  const show = () => {
    setModalDisplay(true)
  }

  return (
    <div action="" className="create-post" onClick={show}>
      <div className="profile-photo">
        <img src="./images/profile-1.jpg" />
      </div>
      <div id="create-post" >What's on your mind, Diana ?</div>
    </div>
  )
}

export default CreatePost
