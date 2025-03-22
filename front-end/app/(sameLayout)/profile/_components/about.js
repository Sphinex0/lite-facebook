import React, { useEffect, useState } from 'react'
import { Cake, Email, Info, PrivacyTip } from '@mui/icons-material'
import styles from "./about.module.css"
import { useRouter } from 'next/navigation'
import { FetchApi } from '@/app/helpers'
const About = ({ user_id, action }) => {
    const [profileInfo, setProfileInfo] = useState({})
    const [privacy, setPrivacy] = useState("")


    const redirect = useRouter()
    const fetchProfileInfo = async (signal) => {
        try {
            const response = await FetchApi("/api/profile/about", redirect, {
                method: "POST",
                body: JSON.stringify({ id: user_id }),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const profileData = await response.json()
                if (profileData) {
                    setProfileInfo(profileData)
                    console.log(profileData)
                    setPrivacy(profileData.privacy)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }
    // /api/profile/update
    const changePrivacy = async (privacy) => {
        try {
            const response = await FetchApi("/api/profile/update",redirect, {
                method: "POST",
                body: JSON.stringify({ privacy }),
            })

            console.log("status:", response.status)
            if (response.ok) {
                const profileData = await response.json()
                console.log(profileData, "response")

                setPrivacy(privacy)
            }

        } catch (error) {
            console.log(error)
        }

    }
    useEffect(() => {
        const controller = new AbortController()
        fetchProfileInfo(controller.signal)
        return () => controller.abort()
    }, [])

    return (
        <div className='feeds'>
            {profileInfo.id &&
                <div className='feed' style={{ display: "flex", justifyContent: "space-evenly", flexWrap: "wrap", minHeight: "250px", alignItems: "center" }}>
                    <div className={styles.infoItem}>
                        {console.log(action)}

                        <span>
                            <select value={privacy} disabled={action == "edit" ? false : true}
                                onChange={(e) => {
                                    changePrivacy(e.target.value)
                                }}
                            >
                                <option value={"public"}>Public</option>
                                <option value={"private"}>Private</option>
                            </select>
                        </span>
                        <PrivacyTip />
                    </div>
                    {profileInfo.about &&
                        <div className={styles.infoItem}>
                            <span>{profileInfo.about}</span>
                            <Info />
                        </div>
                    }
                    <div className={styles.infoItem}>
                        <span>{profileInfo.email}</span>
                        <Email />
                    </div>
                    <div className={styles.infoItem}>
                        <span> {new Date(profileInfo.date_birth).toLocaleDateString()}</span>
                        <Cake />
                    </div>


                </div>
            }

        </div>
    )
}

export default About
