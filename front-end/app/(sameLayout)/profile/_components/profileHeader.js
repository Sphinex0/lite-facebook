"use client"
import { useState } from 'react'
import styles from './profileHeader.module.css'
import FollowBtn from './followBtn'
import UnfollowBtn from './unfollowBtn'
import EditBtn from './editBtn'
import PendingBtn from './pendingBtn'

const ProfileHeader= ({profileInfo, followersCount, followingsCount, profileNav, setProfileNav, actionType}) =>{
    const [action, setAction] = useState(actionType)
    const [followings, setFollowings] = useState(followingsCount)
    const [followers, setFollowers] = useState(followersCount)
  return (
    <div>
      <div className={styles.profileHeader}>
        <div className={styles.basicInfo}>
        <div className={styles.p10}>
          <img className={styles.image} src={`http://localhost:8080/public/pics/${profileInfo?.image || "default-profile.png"}`} />
        </div>
        <div className={styles.g2}>
          
          <h1 className={styles.title}>{profileInfo.first_name} {profileInfo.last_name}</h1>
          { profileInfo.nickname && <span className={styles.nickname}>@{profileInfo.nickname}</span>}<br/>
          <span className={styles.followText}>{followers} followers</span><br/>
          <span className={styles.followText}>{followings} followings</span>
        </div>
        <div className={`${styles.g1} ${styles.btnSection}`}>
            {action == "follow" ? <FollowBtn user_id={profileInfo.id} setAction={setAction} setFollowers={setFollowers}/>:""}
            {action == "unfollow" ? <UnfollowBtn user_id={profileInfo.id} setAction={setAction} setFollowers={setFollowers}/>:""}
            {action == "edit" ? <EditBtn setProfileNav={setProfileNav}/>:""}
            {action == "pending" ?<PendingBtn/>:""}
          
        </div>
        </div>
        <div className={styles.profileNav}>
          <ul className={styles.navUl}>
            <li className={`${styles.navLi} ${profileNav ==="posts"? styles.active:""}`} onClick={()=>setProfileNav("posts")}>Posts</li>
            <li className={`${styles.navLi} ${profileNav ==="about"? styles.active:""}`} onClick={()=>setProfileNav("about")}>About</li>
            <li className={`${styles.navLi} ${profileNav ==="followings"? styles.active:""}`} onClick={()=>setProfileNav("followings")}>Following</li>
            <li className={`${styles.navLi} ${profileNav ==="followers"? styles.active:""}`} onClick={()=>setProfileNav("followers")}>Followers</li>
          </ul>

        </div>
      </div>

    </div>
  )
}
export default ProfileHeader;