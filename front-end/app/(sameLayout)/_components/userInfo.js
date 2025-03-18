import { timeAgo } from '@/app/helpers'
import Link from 'next/link'
import React from 'react'

const UserInfo = ({ userInfo, articleInfo, group , onlineDiv , lastMessage}) => {
    return (
        <Link href={onlineDiv ? "" : `/profile/${userInfo.id}` }>
        <div className="user">
            <div className="profile-wrapper">
                <div className="profile-photo">
                    <img
                        src={`/images/${(userInfo && userInfo.image) || (group && group.image) || "profile-13.jpg"}`}
                        alt="Profile Photo" />
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
                        {articleInfo.parent ==null && <small>{articleInfo.privacy} <strong> .</strong></small>}  <small>{timeAgo(articleInfo.created_at)}</small>
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
