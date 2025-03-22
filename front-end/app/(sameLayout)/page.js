'use client'

import { useEffect, useRef, useState } from "react"
import Post from "./_components/post"
import CreatePost from "./_components/createPost"
import CreatePostModal from "./_components/createPostModal"
import { FetchApi, useOnVisible } from "../helpers"
import PostList from "./_components/postList"
import { redirect } from "next/navigation"
import { useRouter } from "next/navigation"

export default function Posts() {
    const [posts, setPosts] = useState([])
    const [modalDisplay, setModalDisplay] = useState(false)
    
    const redirect = useRouter()

    const lastPostElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now()))
    const fetchData = async (signal) => {
        try {
            console.log("9bl")
            const response = await FetchApi("/api/posts",redirect, {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before: before.current }),
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
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        const controller = new AbortController();
        fetchData(controller.signal)

        return ()=>{
            controller.abort()
        }

    }, [])

    useOnVisible(lastPostElementRef, fetchData)



    return (
        <>
            <CreatePost setModalDisplay={setModalDisplay} />
            {modalDisplay ? <CreatePostModal setModalDisplay={setModalDisplay} setPosts={setPosts} /> : ""}
            <PostList posts={posts} reference={lastPostElementRef} />
        </>
)
}