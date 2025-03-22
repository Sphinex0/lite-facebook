import React from 'react'

const GroupInfo = ({groupInfo}) => {
  return (
    <div className='feed'>
      <div className="user">
          <div className="profile-photo">
              <img src={`./pics/${groupInfo.image}`} />
          </div>
          <div className="ingo">
              <h3>{groupInfo.title} </h3>  
              <small>{groupInfo.description}</small> 
          </div>
      </div>
    </div>
  )
}

export default GroupInfo
