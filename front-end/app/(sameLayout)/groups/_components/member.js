import React from 'react'

const MemberInfo = ({memberInfo}) => {
  return (
    <div className='feed'>
      <div className="user">
          <div className="profile-photo">
              <img src={`http://localhost:8080/public/pics/${memberInfo.image || "default-profile.png"}`}/>
          </div>
          <div className="ingo">
              <h3>{memberInfo.first_name} {memberInfo.last_name} </h3>  
              {/* <small>{memberInfo.description}</small>  */}
          </div>
      </div>
    </div>
  )
}

export default MemberInfo
