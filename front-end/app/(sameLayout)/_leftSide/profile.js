"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import UserInfo from "../_components/userInfo";
import { useWorker } from "@/app/_Context/WorkerContext";

const Profile = () => {
    const {user} = useWorker()
    const [userInfo, setUserInfo] = useState(user)
    useEffect(()=>{
        setUserInfo(user)
    },[user])

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
        {console.log("#######", user)}
        <UserInfo userInfo={userInfo}/>
    </div>
    )
    
}

export default Profile;