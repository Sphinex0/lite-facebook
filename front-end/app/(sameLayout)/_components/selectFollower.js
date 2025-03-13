import React, { useEffect, useRef, useState } from 'react'
import styles from "./selectFollower.module.css"
import UserInfo from './userInfo'
const SelectFollower = () => {
    const [followers, setFollowers] = useState([])
    const before = useRef(Math.floor(Date.now() / 1000))

    const fetchFollowers = async () => {
        try {
            //const before = posts.length > 0 ? posts[posts.length - 1].article.created_at : Math.floor(Date.now() / 1000)
            console.log(before, followers)
            const response = await fetch("http://localhost:8080/api/followers", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before: before.current })
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followersData = await response.json()
                if (followersData) {
                    setFollowers((prv) => [...prv, ...followersData])
                    before.current = followersData[followersData.length-1].modified_at
                }
                console.log(followersData)
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(()=>{
        console.log("hereeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee  ")
        fetchFollowers()
    },[])


  return (
    <div className={styles.container}>
        {followers.map((userInfo)=>{
            return <UserInfo userInfo={userInfo} key={userInfo.id}/>
        })}
    </div>
  )
}

export default SelectFollower
