import React from 'react'
import styles from "./profileHeader.module.css"

const UnfollowBtn = ({user_id, setAction, setFollowers}) => {
    const unfollowUser = async()=>{
        try {
            const response = await fetch("http://localhost:8080/api/follow", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ user_id })
            })
            console.log(JSON.stringify({ user_id }))
            console.log("status:", response.status)
            if (response.ok) {
                setAction("follow")
                setFollowers((prv)=>prv-1)
            }
    
        } catch (error) {
            console.log(error)
        }
    }
  return (
    <button className={`${styles.editProfileBtn}`} onClick={()=>unfollowUser()}>Unfollow</button>

  )
}

export default UnfollowBtn
