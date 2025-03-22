'use client'

import { useState } from "react"
import Posts from "./posts"
import { useRouter } from "next/navigation"
import { FetchApi } from "@/app/helpers"
import { red } from "@mui/material/colors"

const Users = () => {
  const [data , setData] = useState([])
  const redirect = useRouter()



  const fetchPosts = async () => {
    try {
      const before = data.length > 0 ? data[data.length-1].article.created_at : Math.floor(Date.now())
      const response = await FetchApi("/api/posts",redirect,{
        method:"POST",
        body: JSON.stringify({before})
      })

      console.log("status:", response.status)
      if (response.ok){
        const postsData = await response.json()
        if (postsData){
          setData([...data,...postsData])
          console.log(postsData)
        }
      }

    } catch (error) {
      console.log(error)
    }

  }

  const fetchComments = async() => {
    try {
      const response = await FetchApi("/api/comments",redirect,{
        method:"POST",
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
      const response = await FetchApi("/api/followers",redirect,{
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
      const response = await FetchApi("/api/followings",redirect,{
        method:"POST",
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
      const response = await FetchApi("/api/follow/requests",redirect,{
        method:"POST",
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
      const response = await FetchApi("/api/login",redirect,{
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
    <div className="container">
      <button onClick={() => login()}>login</button>
      <button onClick={() => fetchPosts()}>get posts</button>
      <button onClick={() => fetchComments()}>get comments</button>
      {/* <button onClick={() => fetchFollowers()}>get followers</button> */}
      <button onClick={() => fetchFollowings()}>get followings</button>
      <button onClick={() => fetchFollowRequests()}>get follow -requests</button>
      <div>
        {data.length>0 ? <Posts posts={data} /> : "loading"}
      </div>
    </div>
  )
}
export default Users;