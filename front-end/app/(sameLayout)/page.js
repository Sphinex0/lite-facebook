'use client'

import { useEffect, useState } from "react"
import Post from "../_components/post"

export default function Posts() {
    const [posts, setPosts] = useState([])

    const fetchData = async () => {
        try {
            console.log("res")
            const before = posts.length > 0 ? posts[posts.length - 1].article.created_at : Math.floor(Date.now() / 1000)
            const response = await fetch("http://localhost:8080/api/posts", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ before })
            })

            console.log("status:", response.status)
            if (response.ok) {
                const postsData = await response.json()
                if (postsData) {
                    setPosts([...posts, ...postsData])
                    console.log(postsData)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
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
            <form action="" class="create-post" onSubmit='return false'>
                <div class="profile-photo">
                    <img src="./images/profile-1.jpg" />
                </div>
                <input type="text" placeholder="What's on your mind, Diana ?" id="create-post" />
                <input type="submit" value="Post" class="btn btn-primary" />
            </form>
            <div className="feeds" >
                {posts.map((postInfo, index) => {
                    return <Post postInfo={postInfo} key={index} />
                })}
            </div>
        </>

    )
}