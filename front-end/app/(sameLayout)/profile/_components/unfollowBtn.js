import React from 'react'
import styles from "./profileHeader.module.css"
import { useRouter } from 'next/navigation'
import { FetchApi } from '@/app/helpers'

const UnfollowBtn = ({user_id, setAction, setFollowers}) => {
    const redirect = useRouter()
    const unfollowUser = async()=>{
        try {
            const response = await FetchApi("/api/follow", redirect,{
                method: "POST",
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
