import React, { useEffect, useRef, useState } from 'react'
import UserInfo from '../../_components/userInfo'
import Link from 'next/link'
import { FetchApi, useOnVisible } from '@/app/helpers'
import { useRouter } from 'next/navigation'

const Followers = ({ user_id, setIsAllowed }) => {
    const [followers, setFollowers] = useState([])
    const before = useRef(Math.floor(Date.now()))
    const lastElementRef = useRef(null)

    const redirect = useRouter()

    const fetchFollowers = async (signal) => {
        try {
            const response = await FetchApi("/api/followers", redirect, {
                method: "POST",
                body: JSON.stringify({ before: before.current, user_id }),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followersData = await response.json()
                if (followersData) {
                    setFollowers((prv) => [...prv, ...followersData])
                    before.current = followersData[followersData.length - 1].modified_at
                    console.log(followersData)
                    setIsAllowed(true)
                }

            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        const controller = new AbortController()
        fetchFollowers(controller.signal)
        return () => controller.abort()
    }, [])
    useOnVisible(lastElementRef, fetchFollowers)
    return (
        <div className='feeds'>
            {followers.map((userInfo, index) => {
                if (index == followers.length - 1) {
                    return <div className='feed' key={`user${userInfo.id}`} ref={lastElementRef}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>
                }

                return <div className='feed' key={`user${userInfo.id}`}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>

            })}
        </div>
    )
}

export default Followers
