async function Checkuservalidity() {
    try {
        const response = await fetch("/checkuser")
        return response.json()
    }catch(err){
        // redirecte to intenalserver err
    }
}