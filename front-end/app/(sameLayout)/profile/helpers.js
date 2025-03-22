import { FetchApi } from "@/app/helpers"
import { red } from "@mui/material/colors"

export const fetchProfile = async(setProfileInfo, id,redirect)=>{
    try {
        const response = await FetchApi("/api/profile",redirect, {
            method: "POST",
            body: JSON.stringify({ id })
        })
        console.log(JSON.stringify({ id }))
        console.log("status:", response.status)
        if (response.ok) {
            const profileData = await response.json()
            setProfileInfo(profileData)
        }

    } catch (error) {
        console.log(error)
    }
}