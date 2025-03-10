'use client'
const Users = () => {
  const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/posts",{
        method:"POST",
        credentials: "include",
        body:"{}"
      })
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }
  }

  const fetchComments = async() => {
    try {
      const response = await fetch("http://localhost:8080/api/comments",{
        method:"POST",
        credentials: "include",
        body:JSON.stringify({"parent":1})
      })
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }

  }

  const fetchFollowers = async() => {
    try {
      const response = await fetch("http://localhost:8080/api/followers",{
        method:"POST",
        credentials: "include",
        body:"{}"
      })
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }

  }

  const fetchFollowings = async() => {
    try {
      const response = await fetch("http://localhost:8080/api/followings",{
        method:"POST",
        credentials: "include",
        body:"{}"
      })
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }

  }
  const fetchFollowRequests = async() => {
    try {
      const response = await fetch("http://localhost:8080/api/follow/requests",{
        method:"POST",
        credentials: "include",
        body:"{}"
      })
      console.log("status:", response.status)
      if (response.ok){
        const data = await response.json()
        console.log(data)
      }

    } catch (error) {
      console.log(error)
    }

  }

  const login = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/login",{
        credentials: "include"
      })
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
      <button onClick={() => fetchComments()}>get comments</button>
      <button onClick={() => fetchFollowers()}>get followers</button>
      <button onClick={() => fetchFollowings()}>get followings</button>
      <button onClick={() => fetchFollowRequests()}>get follow requests</button>
    </div>
  )
}
export default Users;