import React, { useEffect, useRef, useState } from 'react'
import PostList from '../../_components/postList'
import { FetchApi, useOnVisible } from '@/app/helpers'
import { useRouter } from 'next/navigation'


const Posts = ({user_id, setIsAllowed}) => {
    const [posts, setPosts] = useState([])
    
    const redirect = useRouter()
    const lastPostElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now()))

        const fetchProfilePosts = async (signal) => {
            try {
                const response = await FetchApi("/api/profile/posts",redirect, {
                    method: "POST",
                    body: JSON.stringify({ before: before.current, user_id }),
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
            fetchProfilePosts(controller.signal)
    
            return ()=>{
                controller.abort()
            }
    
        }, [])
    
        useOnVisible(lastPostElementRef, fetchProfilePosts)
  return (
    <PostList posts={posts} reference={lastPostElementRef} />
  )
}

export default Posts
