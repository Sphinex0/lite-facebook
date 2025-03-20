import React from 'react'
import styles from "./profileHeader.module.css"

const FollowBtn = ({user_id, setAction, setFollowers}) => {
    const followUser = async()=>{
        try {
            const response = await fetch("http://localhost:8080/api/follow", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ user_id })
            })
            console.log(JSON.stringify({ user_id }))
            console.log("status:", response.status)
            if (response.ok) {
                const inviteData = await response.json()
                if (inviteData.status == "accepted"){

                    setAction("unfollow")
                    setFollowers((prv)=>prv+1)
                }else if(inviteData.status == "pending" ){
                    setAction("pending")
                }
            }
    
        } catch (error) {
            console.log(error)
        }
    }
  return (
    <button className={`${styles.editProfileBtn} ${styles.purple}`} onClick={()=>followUser()}>Follow</button>
  )
}

export default FollowBtn
