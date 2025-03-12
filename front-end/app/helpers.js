async function Checkuservalidity() {
    try {
        const response = await fetch("/checkuser")
        return response.json()
    }catch(err){
       return false
    }
}