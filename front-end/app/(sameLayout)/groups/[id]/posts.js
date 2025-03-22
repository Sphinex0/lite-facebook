import React, { useEffect, useRef, useState } from 'react'
import PostList from '../../_components/postList'
import { FetchApi, useOnVisible } from '@/app/helpers'
import CreatePost from '../../_components/createPost'
import CreatePostModal from '../../_components/createPostModal'
import { useRouter } from 'next/navigation'

const Posts = ({ groupID, setIsAllowed }) => {
    const [posts, setPosts] = useState([])
    const [modalDisplay, setModalDisplay] = useState(false)

    const redirect = useRouter()

    const lastPostElementRef = useRef(null)
    const before = useRef(Math.floor(Date.now()))

    const fetchGroupPosts = async (signal) => {
        try {
            const response = await FetchApi("/api/group/posts",redirect, {
                method: "POST",
                body: JSON.stringify({ before: before.current, group_id: +groupID }),
                signal
            })

            console.log("status:", response.status)
            if (response.ok) {
                const postsData = await response.json()
                if (postsData) {
                    setPosts((prv) => [...prv, ...postsData])
                    before.current = postsData[postsData.length - 1].article.created_at
                    console.log("last created at", postsData[postsData.length - 1].article.created_at)
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

        return () => {
            controller.abort()
        }

    }, [])

    useOnVisible(lastPostElementRef, fetchGroupPosts)
    return (<>
        <CreatePost setModalDisplay={setModalDisplay} />
        {modalDisplay ? <CreatePostModal setModalDisplay={setModalDisplay} setPosts={setPosts} group={groupID} /> : ""}
        <PostList posts={posts} reference={lastPostElementRef} />
    </>
    )
}

export default Posts
