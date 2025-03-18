import React, { useEffect,  useState } from 'react'
import { Cake, Email, Info, PrivacyTip  } from '@mui/icons-material'
import styles from "./about.module.css"
const About = ({user_id}) => {
        const [profileInfo, setProfileInfo] = useState({})
    
        const fetchProfileInfo = async (signal) => {
            try {
                const response = await fetch("http://localhost:8080/api/profile/about", {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({id:user_id }),
                    signal
                })
    
                console.log("status:", response.status)
                if (response.ok) {
                    const profileData = await response.json()
                    if (profileData){
                        setProfileInfo(profileData)
                        console.log(profileData)
                    }
                }
    
            } catch (error) {
                console.log(error)
            }
    
        }
    
        useEffect(() => {
            const controller = new AbortController()
            fetchProfileInfo(controller.signal)
            return ()=>controller.abort()
        }, [])

  return (
    <div className='feeds'>
        {profileInfo.id &&  
        <div className='feed' style={{display:"flex", justifyContent:"space-evenly",flexWrap:"wrap",  minHeight:"250px", alignItems:"center"}}>
            <div  className={styles.infoItem}>
               <span> {profileInfo.privacy}</span>
                <PrivacyTip/>
            </div>
            {profileInfo.about &&
            <div className={styles.infoItem}>
                <span>{profileInfo.about}</span>
                <Info/>
            </div>
            }
            <div className={styles.infoItem}>
                <span>{profileInfo.email}</span>
                <Email/>
            </div>
            <div  className={styles.infoItem}>
               <span> {new Date(profileInfo.date_birth*1000).toLocaleDateString()}</span>
                <Cake/>
            </div>
            

        </div>
        }
       
    </div>
  )
}

export default About
