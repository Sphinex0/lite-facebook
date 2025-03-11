import Post from "./post"

export default function Posts({posts}) {
    return (
        <div className="feeds">
            {posts.map((postInfo,index)=>{
                return <Post postInfo={postInfo} key={index}/>
            })}
        </div>
    )
}