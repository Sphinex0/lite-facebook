"use client"
import { timeAgo } from '@/app/helpers'
import Image from 'next/image'
import Link from 'next/link'
import React, { useEffect, useState } from 'react'

const UserInfo = ({redirect, userInfo, articleInfo, group, onlineDiv, lastMessage }) => {

    const [imageSrc, setImageSrc] = useState(`http://localhost:8080/public/pics/${userInfo ? userInfo.image : (group && group.image)}`)

    if (userInfo == null) { 
        userInfo = {}
    }

    useEffect(() => {
        setImageSrc(`http://localhost:8080/public/pics/${userInfo?.first_name ? userInfo.image : (group && group.image)}`)
    },[userInfo])


    return (
        <Link href={onlineDiv || typeof redirect == "boolean" ? "" : `/profile/${userInfo.id}`}>
            <div className="user">
                <div className="profile-wrapper">
                    <div className="profile-photo">
                        {/* <Image
                            src={imageSrc}
                            alt={"ess"}
                            width={50} 
                            height={50}
                            onError={() => setImageSrc('http://localhost:8080/public/pics/default-profile.png')}
                        /> */}
                        <img
                        src={`http://localhost:8080/public/pics/${userInfo.image ? userInfo.image : (group && group.image)}`}
                        alt="Profile Photo" />
                    </div>
                    {onlineDiv && <div className={userInfo.first_name && `status ${userInfo && (userInfo.online ? "online" : "offline")}`}></div>}
                </div>
                <div className="ingo">
                    <h3>
                        {
                            (group && group.title) || userInfo && `${userInfo.first_name} ${userInfo.last_name}`
                        }
                    </h3>
                    {articleInfo &&
                        <>
                            {articleInfo.parent == null && <small>{articleInfo.privacy} <strong> .</strong></small>}  <small>{timeAgo(articleInfo.created_at)}</small>
                        </>
                    }
                    {
                        lastMessage && <>
                            <small>{lastMessage}</small>
                        </>
                    }
                </div>
            </div>
        </Link>
    )
}

export default UserInfo
