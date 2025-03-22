import React, { useEffect, useRef, useState } from 'react'
import UserInfo from '../../_components/userInfo'
import { FetchApi, useOnVisible } from '@/app/helpers'
import { useRouter } from 'next/navigation'

const Followings = ({ user_id }) => {
    const [followings, setFollowings] = useState([])
    const before = useRef(Math.floor(Date.now()))
    const lastElementRef = useRef(null)

    const redirect = useRouter()

    const fetchFollowings = async (signal) => {
        try {
            const response = await FetchApi("/api/followings", redirect, {
                method: "POST",
                body: JSON.stringify({ before: before.current, user_id }),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followingsData = await response.json()
                if (followingsData) {
                    setFollowings((prv) => [...prv, ...followingsData])
                    before.current = followingsData[followingsData.length - 1].modified_at
                    console.log(followingsData)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        const controller = new AbortController()
        fetchFollowings(controller.signal)
        return () => controller.abort()
    }, [])
    useOnVisible(lastElementRef, fetchFollowings)
    return (
        <div className='feeds' style={{ display: "flex", flexWrap: "wrap" }}>
            {followings.map((userInfo, index) => {
                if (index == followings.length - 1) {
                    return <div className='feed' key={`user${userInfo.id}`} ref={lastElementRef}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>
                }
                return <div className='feed' key={`user${userInfo.id}`}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>
            })}
        </div>
    )
}

export default Followings
