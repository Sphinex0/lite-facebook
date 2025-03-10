'use client'
const Users = () => {
  const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/posts")
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }
  }

  const fetchComments = () => {

  }

  const fetchFollowers = () => {

  }

  const fetchFollowings = () => {

  }

  const followRequests = () => {

  }

  const login = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/login")
      console.log("status:", response.status)
      console.log(response)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }
  }


  return (
    <div>
      <button onClick={() => login()}>login</button>
      <button onClick={() => fetchPosts()}>get posts</button>
    </div>
  )
}
export default Users;