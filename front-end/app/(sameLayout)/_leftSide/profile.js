"use client";

import Link from "next/link";
import { useState } from "react";
import UserInfo from "../_components/userInfo";

const Profile = () => {
    const [userInfo, setUserInfo] = useState(JSON.parse(localStorage.getItem("user")))

    return (
        <div  className="profile">
        {/* <div className="profile-photo">
            <img src="/images/profile-1.jpg"/>
        </div>
        <div className="handle">
            <h4>Nia Ridania</h4>
            <p className="text-muted">
                @niaridania
            </p>
        </div> */}
        <UserInfo userInfo={userInfo}/>
    </div>
    )
    
}

export default Profile;