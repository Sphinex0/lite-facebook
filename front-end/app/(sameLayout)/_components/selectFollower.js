import React, { useEffect, useRef, useState } from 'react'
import styles from "./selectFollower.module.css"
import UserInfo from './userInfo'
import { useOnVisible } from '@/app/helpers'
const SelectFollower = () => {
    const [followers, setFollowers] = useState([])
    const container = useRef(null)
    const before = useRef(Math.floor(Date.now() / 1000))

    const fetchFollowers = async (first = false) => {
        const lastElementRef = useRef(null)
        try {

            const response = await fetch("http://localhost:8080/api/followers", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before: before.current })
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followersData = await response.json()
                if (followersData != followers) {
                    if (first) {
                        setFollowers(followersData)
                    } else {
                        setFollowers((prv) => [...prv, ...followersData])
                    }
                    before.current = followersData[followersData.length - 1].modified_at
                }
                console.log(followersData)
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        fetchFollowers(true)
    }, [])
    useOnVisible(lastElementRef, fetchFollowers)


    return (<>
        <h3>choose who can see your post:</h3>
        <div className={styles.container} ref={container}>
            {followers.map((userInfo) => {
                return <div ref={lastElementRef} className={styles.fullUser} key={`user${userInfo.id}`}><label htmlFor={`user${userInfo.id}`}><UserInfo userInfo={userInfo} key={userInfo.id} /></label> <input type='checkbox' id={`user${userInfo.id}`} name='users' value={userInfo.id} /></div>
            })}
        </div>
    </>

    )
}

export default SelectFollower
