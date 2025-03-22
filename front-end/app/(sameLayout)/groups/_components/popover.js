import React, { useEffect, useRef, useState } from 'react'
import styles from "./popover.module.css"
import UserInfo from '../../_components/userInfo'
import { Add } from '@mui/icons-material'
import JoinGroup from '../[id]/function'
import { useOnVisible } from '@/app/helpers'



const Popover = ({group_id}) => {
    const [followers, setFollowers] = useState([])
    const before = useRef(0)
    const lastElementRef = useRef(null)

    const fetchFollowers = async (signal) => {
        try {
            const response = await fetch("http://localhost:8080/api/group/invitelist", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before: before.current , group_id}),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followersData = await response.json()
                if (followersData){
                    setFollowers((prv) => [...prv, ...followersData])
                    before.current = followersData[followersData.length - 1].id
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


    return (<>
        <div className={styles.container} >
            {followers.map((userInfo, index) => {
                if (index == followers.length-1){
                    return <div ref={lastElementRef} className={styles.fullUser} key={`user${userInfo.id}`}>
                        <label htmlFor={`user${userInfo.id}`}>
                            <UserInfo redirect={false} userInfo={userInfo} key={userInfo.id} />
                            </label>                     
                            <button className={styles.addMember} onClick={()=>{
                        JoinGroup(id, groupData.group_info.creator, setIsAction, isAction)
      
                      }}><Add/></button>
                      </div>
                }
                return <div className={styles.fullUser} key={`user${userInfo.id}`}>
                    <label htmlFor={`user${userInfo.id}`}>
                        <UserInfo redirect={false} userInfo={userInfo} key={userInfo.id} />
                    </label> 
                    <button className={styles.addMember} onClick={()=>{
                  JoinGroup(id, groupData.group_info.creator, setIsAction, isAction)
                }}><Add/></button></div>
            })}
        </div>
    </>

    )
}

export default Popover
