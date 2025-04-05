import React, { useEffect, useRef, useState } from 'react'
import PostList from '../../_components/postList'
import { FetchApi, useOnVisible } from '@/app/helpers'
import MemberInfo from '../_components/member'
import { useRouter } from 'next/navigation'
import UserInfo from '../../_components/userInfo'

const Members = ({ groupID }) => {
    const [members, setMembers] = useState([])

    const redirect = useRouter()

    const lastElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now()))

    const fetchGroupMembers = async (signal) => {
        try {
            const response = await FetchApi("/api/invites/members", redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ group_id: parseInt(groupID) }),
                signal

            })

            if (response.ok) {
                const membersData = await response.json()
                if (membersData) {
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

        return () => {
            controller.abort()
        }

    }, [])

    useOnVisible(lastElementRef, fetchGroupMembers)
    return (
        <div className='feeds'>
            {members.map((memberInfo, index) => {
                if (index == members.length - 1) {
                    return <div className='feed' key={`member${memberInfo.id}`}>
                        <UserInfo  userInfo={memberInfo} reference={lastElementRef} />
                    </div>

                }
                return <div className='feed' key={`member${memberInfo.id}`}>
                    <UserInfo  userInfo={memberInfo} />
                </div>


            })}
        </div>
    )
}

export default Members
