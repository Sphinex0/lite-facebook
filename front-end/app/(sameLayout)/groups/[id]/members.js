import React, { useEffect, useRef, useState } from 'react'
import PostList from '../../_components/postList'
import { useOnVisible } from '@/app/helpers'
import MemberInfo from '../_components/member'

const Members = ({groupID}) => {
    const [members, setMembers] = useState([])
    
    const lastElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now() / 1000))

        const fetchGroupMembers = async (signal) => {
            try {
                const response = await fetch("http://localhost:8080/api/invites/members", {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({ before: before.current, group_id:+groupID }),
                    signal
                    
                })
    
                console.log("status:", response.status)
                if (response.ok) {
                    const membersData = await response.json()
                    if (membersData) {
                        console.log(membersData)
                        setMembers((prv) => [...prv, ...membersData])
                        //before.current = membersData[membersData.length-1].article.created_at
                    }
                    
                }

    
            } catch (error) {
                console.log(error)
            }
    
        }
    
        useEffect(() => {
            const controller = new AbortController();
            fetchGroupMembers(controller.signal)
    
            return ()=>{
                controller.abort()
            }
    
        }, [])
    
        useOnVisible(lastElementRef, fetchGroupMembers)
  return (
    <div className='feeds'>
        {members.map((memberInfo, index)=>{
            if (index == members.length-1){
                return  <MemberInfo memberInfo={memberInfo} reference={lastElementRef} />

            }
            return   <MemberInfo memberInfo={memberInfo} reference={lastElementRef} />


        })}
    </div>
  )
}

export default Members
