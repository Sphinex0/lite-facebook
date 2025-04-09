import React, { useEffect, useRef, useState } from 'react'
import styles from "./popover.module.css"
import UserInfo from '../../_components/userInfo'
import { Add, Done } from '@mui/icons-material'
import JoinGroup from '../[id]/function'
import { FetchApi, useOnVisible } from '@/app/helpers'
import { useRouter } from 'next/navigation'



const Popover = ({ group_id }) => {
    const [followers, setFollowers] = useState([])
    const before = useRef(0)
    const lastElementRef = useRef(null)
    const redirect = useRouter()
    const [clicked, setClicked] = useState([])

    const fetchFollowers = async (signal) => {
        try {
            const response = await FetchApi("/api/group/invitelist", redirect, {
                method: "POST",
                body: JSON.stringify({ before_id: before.current, group_id }),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const followersData = await response.json()
                if (followersData) {
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
        return () => controller.abort()
    }, [])
    useOnVisible(lastElementRef, fetchFollowers)


    return (<>
        <div className={styles.container} >
            {followers.map((userInfo, index) => {
                if (index == followers.length - 1) {
                    return <div ref={lastElementRef} className={styles.fullUser} key={`user${userInfo.id}`}>
                        <label htmlFor={`user${userInfo.id}`}>
                            <UserInfo redirect={false} userInfo={userInfo} key={userInfo.id} />
                        </label>
                        <button className={styles.addMember} onClick={(e) => {
                            JoinGroup(group_id, userInfo.id)
                            setClicked(prv => [...prv, userInfo.id])

                        }}>{clicked.indexOf(userInfo.id) != -1 ? <Done /> : <Add />}</button>
                    </div>
                }
                return <div className={styles.fullUser} key={`user${userInfo.id}`}>
                    <label htmlFor={`user${userInfo.id}`}>
                        <UserInfo redirect={false} userInfo={userInfo} key={userInfo.id} />
                    </label>
                    <button className={styles.addMember} onClick={() => {
                        JoinGroup(group_id, userInfo.id)
                        setClicked(prv => [...prv, userInfo.id])

                    }}>{clicked.indexOf(userInfo.id) != -1 ? <Done /> : <Add />}</button>
                </div>
            })}
        </div>
    </>

    )
}

export default Popover
