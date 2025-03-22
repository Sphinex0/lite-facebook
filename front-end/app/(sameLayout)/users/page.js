"use client"
import React, { useEffect, useRef, useState } from 'react'
import UserInfo from '../_components/userInfo'
import { FetchApi, useOnVisible } from '@/app/helpers'
import { useRouter } from 'next/navigation'

const Page = () => {
    const [users, setUsers] = useState([])
    const redirect = useRouter()

    const lastElementRef = useRef(null)
    const before = useRef(null)
    const fetchData = async (signal) => {
        try {
            console.log(before, users)
            const response = await FetchApi("/api/users",redirect, {
                method: "POST",
                body: JSON.stringify({ before: before.current }),
                signal
                
            })

            console.log("status:", response.status)
            if (response.ok) {
                const usersData = await response.json()
                if (usersData) {
                    setUsers((prv) => [...prv, ...usersData])
                    before.current = usersData[usersData.length-1].id
                    console.log("last created at", usersData[usersData.length-1].id)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        const controller = new AbortController();
        fetchData(controller.signal)

        return ()=>{
            controller.abort()
        }

    }, [])

    useOnVisible(lastElementRef, fetchData)



    return (
        <div className='feeds'>
            {users.map((userInfo , index)=>{
                if (index == users.length-1){
                    return <div className='feed' key={`users${userInfo.id}`} ref={lastElementRef}><UserInfo userInfo={userInfo}/></div>
                }
                return <div className='feed' key={`users${userInfo.id}`} ref={lastElementRef}><UserInfo userInfo={userInfo}/></div>
            })}
        </div>
)
}

export default Page
