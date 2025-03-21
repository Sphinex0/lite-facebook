"use client"
import { timeAgo } from '@/app/helpers'
import Image from 'next/image'
import Link from 'next/link'
import React, { useEffect, useState } from 'react'

const UserInfo = ({ userInfo, articleInfo, group, onlineDiv, lastMessage }) => {

    const [imageSrc, setImageSrc] = useState(`/pics/${userInfo ? userInfo.image : (group && group.image)}`)

    if (userInfo == null) { 
        userInfo = {}
    }

    useEffect(() => {
        setImageSrc(`/pics/${userInfo ? userInfo.image : (group && group.image)}`)
    },[userInfo])


    return (
        <Link href={onlineDiv ? "" : `/profile/${userInfo.id}`}>
            {console.log(userInfo, "inside")}
            <div className="user">
                <div className="profile-wrapper">
                    <div className="profile-photo">
                        {console.log(" imageSrc => ",imageSrc)}
                        <Image
                            src={imageSrc}
                            alt={"ess"}
                            width={50} // Required by next/image
                            height={50} // Required by next/image
                            onError={() => setImageSrc('/pics/default-profile.png')}
                        />
                        {/* <img
                        src={`/pics/${userInfo ? userInfo.image : (group && group.image)}`}
                        alt="Profile Photo" /> */}
                    </div>
                    {onlineDiv && <div className={`status ${userInfo && (userInfo.online ? "online" : "offline")}`}></div>}
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
