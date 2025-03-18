"use client";

import Link from "next/link";

const Profile = () => {

    return (
        <Link  href="/profile" className="profile">
        <div className="profile-photo">
            <img src="/images/profile-1.jpg"/>
        </div>
        <div className="handle">
            <h4>Nia Ridania</h4>
            <p className="text-muted">
                @niaridania
            </p>
        </div>
    </Link>
    )
}

export default Profile;