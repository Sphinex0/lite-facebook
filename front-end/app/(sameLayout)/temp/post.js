import { FetchApi } from "@/app/helpers"
import { useRouter } from "next/navigation"
import { use, useState } from "react"
//import styles from './posts.module.css'

export default function Post({ postInfo }) {
    const [likes, setLikes] = useState(postInfo.likes)
    const [disLikes, setDislikes] = useState(postInfo.disLikes)
    const [likeState, setLikeState] = useState(postInfo.like) //0 1 -1
    const redirect = useRouter()

    const likePost = async (like, article_id) => {
        try {
            const response = await FetchApi("/api/reactions/store",redirect, {
                method: "POST",
                body: JSON.stringify({ like, article_id })
            })

            console.log("status:", response.status)
            if (response.ok) {
                console.log("liked or dislike")
                if (like == 1 && likeState == 1) {
                    setLikes(likes - 1)
                    setLikeState(0)
                } else if (like == -1 && likeState == -1) {
                    setDislikes(disLikes - 1)
                    setLikeState(0)
                } else if (like == 1 && likeState == -1) {
                    setDislikes(disLikes - 1)
                    setLikes(likes + 1)
                    setLikeState(1)
                } else if (like == -1 && likeState == 1) {
                    setDislikes(disLikes + 1)
                    setLikes(likes - 1)
                    setLikeState(-1)
                } else if (like == 1 && likeState == 0) {
                    setLikes(likes + 1)
                    setLikeState(1)
                } else if (like == -1 && likeState == 0) {
                    setDislikes(disLikes + 1)
                    setLikeState(-1)
                }
            }

        } catch (error) {
            console.log(error)
        }

    }

    /**
     *             <hr/>
          <h1>content : {postInfo.article.content}</h1>
          <h3><span onClick={()=>{likePost(1, postInfo.article.id)}}>likes</span>: {likes} || <span onClick={()=>{likePost(-1, postInfo.article.id)}}>dislikes</span>: {disLikes}</h3>
          <hr/>
     */
    return (
        <div className="feed">
            <div className="head">
                <div className="user">
                    <div className="profile-photo">
                        <img src="./images/profile-13.jpg" />
                    </div>
                    <div className="ingo">
                        <h3>lana rose</h3>
                        <small>Morocco, 15 minutes ago</small>
                    </div>
                </div>
            </div>
            <div className="photo">
                <img src="./images/feed-1.jpg" />
            </div>
            <div className="action-button">
                <div className="action-buttons">
                    <span><i className="fa-regular fa-heart"></i></span>
                    <span><i className="fa-regular fa-comment"></i></span>
                </div>
            </div>
        </div>
    )
}