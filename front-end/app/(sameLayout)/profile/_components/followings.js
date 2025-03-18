import React, { useEffect, useRef, useState } from 'react'
import UserInfo from '../../_components/userInfo'
import Link from 'next/link'
import { useOnVisible } from '@/app/helpers'

const Followings = ({user_id}) => {
        const [followings, setFollowings] = useState([])
        const before = useRef(Math.floor(Date.now() / 1000))
        const lastElementRef = useRef(null)
    
        const fetchFollowings = async (signal) => {
            try {
                const response = await fetch("http://localhost:8080/api/followings", {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({ before: before.current,user_id }),
                    signal
                })
    
                console.log("status:", response.status)
                if (response.ok) {
                    const followingsData = await response.json()
                    if (followingsData){
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
            return ()=>controller.abort()
        }, [])
        useOnVisible(lastElementRef, fetchFollowings)
  return (
    <div className='feeds' style={{display:"flex", flexWrap:"wrap"}}>
        {followings.map((userInfo, index) => {
                if (index == followings.length-1){
                    return <div className='feed' key={`user${userInfo.id}`}><Link href={`/profile/${userInfo.id}`} ref={lastElementRef} ><UserInfo userInfo={userInfo} key={userInfo.id} /></Link></div>
                }
                return <div className='feed'  key={`user${userInfo.id}`}><Link href={`/profile/${userInfo.id}`}><UserInfo userInfo={userInfo} key={userInfo.id} /></Link></div>
            })}
    </div>
  )
}

export default Followings
