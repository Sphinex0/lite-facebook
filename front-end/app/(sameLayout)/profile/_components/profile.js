"use client"
import { useEffect, useState } from 'react';
import Followers from './followers';
import Followings from './followings';
import About from './about';
import ProfileHeader from './profileHeader';
import { fetchProfile } from '../helpers';
import Posts from './posts';
import { red } from '@mui/material/colors';
import { useRouter } from 'next/navigation';
const Profile = ({ userID }) => {
  const [profileInfo, setProfileInfo] = useState({})
  const [profileNav, setProfileNav] = useState("posts")
  const [isAllowed, setIsAllowed] = useState(false)

  const redirect = useRouter()



  useEffect(() => {
    fetchProfile(setProfileInfo, userID, redirect)
  }, [])


  return (
    <div>

      {profileInfo.user_info && (
        <ProfileHeader
          profileInfo={profileInfo.user_info}
          followersCount={profileInfo.followers}
          followingsCount={profileInfo.followings}
          actionType={profileInfo.action}
          profileNav={profileNav}
          setProfileNav={setProfileNav}
        />
      )}

      {profileNav == "posts" ? <Posts user_id={userID} setIsAllowed={setIsAllowed} /> : ""}
      {profileNav == "followers" ? <Followers user_id={userID} /> : ""}
      {profileNav == "followings" ? <Followings user_id={userID} /> : ""}
      {profileNav == "about" ? <About user_id={userID} action={profileInfo.action} /> : ""}

      {!isAllowed && <div>join / follow to see</div>}
    </div>
  )
}
export default Profile;