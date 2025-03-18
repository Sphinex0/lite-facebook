import React, { useEffect, useRef, useState } from 'react'
import UserInfo from '../../_components/userInfo'
import Link from 'next/link'
import { useOnVisible } from '@/app/helpers'

const Followers = ({user_id}) => {
        const [followers, setFollowers] = useState([])
        const before = useRef(Math.floor(Date.now() / 1000))
        const lastElementRef = useRef(null)
    
        const fetchFollowers = async (signal) => {
            try {
                const response = await fetch("http://localhost:8080/api/followers", {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({ before: before.current , user_id}),
                    signal
                })
    
                console.log("status:", response.status)
                if (response.ok) {
                    const followersData = await response.json()
                    if (followersData){
                        setFollowers((prv) => [...prv, ...followersData])
                        before.current = followersData[followersData.length - 1].modified_at
                        console.log(followersData)
                    }
                }
    
            } catch (error) {
                console.log(error)
            }
    
        }
    
        useEffect(() => {
            const controller = new AbortController()
            fetchFollowers(controller.signal)
            return ()=>controller.abort()
        }, [])
        useOnVisible(lastElementRef, fetchFollowers)
  return (
    <div className='feeds'>
        {followers.map((userInfo, index) => {
                if (index == followers.length-1){
                    return <div className='feed'  key={`user${userInfo.id}`} ref={lastElementRef}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>
                }
                
                return <div className='feed' key={`user${userInfo.id}`}><UserInfo userInfo={userInfo} key={userInfo.id} /></div>
                
            })}
    </div>
  )
}

export default Followers
