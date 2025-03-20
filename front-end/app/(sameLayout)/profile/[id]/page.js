import React, { use } from 'react'
import ProfileHeader from '../_components/profileHeader'
import Profile from '../_components/profile';

const Page = ({params}) => {
      const id = use(params).id;
    
  return (
    <Profile userID={+id}/>
  )
}

export default Page
