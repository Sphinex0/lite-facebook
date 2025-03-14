'use client'

import { useEffect, useRef, useState } from "react"
import Post from "./_components/post"
import CreatePost from "./_components/createPost"
import CreatePostModal from "./_components/createPostModal"

export default function Posts() {
    const [posts, setPosts] = useState([])
    const [modalDisplay, setModalDisplay] = useState(false)
    const before = useRef(Math.floor(Date.now() / 1000))

    const fetchData = async () => {
        console.log("Posts rendering");
        try {
            //const before = posts.length > 0 ? posts[posts.length - 1].article.created_at : Math.floor(Date.now() / 1000)
            console.log(before, posts)
            const response = await fetch("http://localhost:8080/api/posts", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before: before.current })
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
        console.log("######################################################################################")
        fetchData()
        window.onscroll = () => {
            console.log("here")
            if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
                fetchData()
            }
        }
    }, [])





    return (
        <>
            <CreatePost setModalDisplay={setModalDisplay} />
            {modalDisplay ? <CreatePostModal setModalDisplay={setModalDisplay} setPosts={setPosts} /> : ""}
            <div className="feeds" >
                {console.log("all posts", posts)}
                {posts.map((postInfo) => {
                    return <Post postInfo={postInfo} key={postInfo.article.id} />
                })}
            </div>
        </>

    )
}