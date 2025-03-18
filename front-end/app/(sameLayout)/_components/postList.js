import React from 'react'
import Post from './post'

const PostList = ({posts, reference}) => {
  return (
    <div className="feeds" >
                    {posts.map((postInfo, index) => {
                        if (index == posts.length - 1) {
                            return <Post postInfo={postInfo} key={postInfo.article.id} reference={reference} />
                        }
                        return <Post postInfo={postInfo} key={postInfo.article.id} />
                    })}
                </div>
  )
}

export default PostList
