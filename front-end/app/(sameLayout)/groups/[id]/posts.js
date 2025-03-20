import React, { useEffect, useRef, useState } from 'react'
import PostList from '../../_components/postList'
import { useOnVisible } from '@/app/helpers'

const Posts = ({groupID, setIsAllowed}) => {
    const [posts, setPosts] = useState([])
    
    const lastPostElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now() / 1000))

        const fetchGroupPosts = async (signal) => {
            try {
                const response = await fetch("http://localhost:8080/api/group/posts", {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({ before: before.current, group_id:+groupID }),
                    signal
                    
                })
    
                console.log("status:", response.status)
                if (response.ok) {
                    const postsData = await response.json()
                    if (postsData) {
                        setPosts((prv) => [...prv, ...postsData])
                        before.current = postsData[postsData.length-1].article.created_at
                        console.log("last created at", postsData[postsData.length-1].article.created_at)
                    }
                    setIsAllowed(true)
                }

    
            } catch (error) {
                console.log(error)
            }
    
        }
    
        useEffect(() => {
            const controller = new AbortController();
            fetchGroupPosts(controller.signal)
    
            return ()=>{
                controller.abort()
            }
    
        }, [])
    
        useOnVisible(lastPostElementRef, fetchGroupPosts)
  return (
    <PostList posts={posts} reference={lastPostElementRef} />
  )
}

export default Posts
