import React from 'react'

const UserInfo = ({ userInfo, articleInfo , group }) => {
    return (
        <div className="user">
            <div className="profile-photo">
                <img src={`./images/${userInfo.image || group.image || "profile-13.jpg"}`} />
            </div>
            <div className="ingo">
                <h3>
                    {
                        group.title || `${userInfo.first_name} ${userInfo.last_name}`
                    }
                </h3>
                {articleInfo && <small>{timeAgo(articleInfo.created_at)}</small>}
            </div>
        </div>
    )
}

export default UserInfo
