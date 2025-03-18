export const fetchProfile = async(setProfileInfo, id)=>{
    try {
        const response = await fetch("http://localhost:8080/api/profile", {
            method: "POST",
            credentials: "include",
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