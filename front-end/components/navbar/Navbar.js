'use client'
import './navbar.css'
import ChatPopover from '@/components/popovers/chatpopover/page'
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupOutlinedIcon from '@mui/icons-material/GroupOutlined'
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined'
import MailOutlinedIcon from '@mui/icons-material/MailOutlined'
import { useState } from 'react'

export default function Navbar () {
  const [bool, setbool] = useState(false)
  function handleclick () {
    setbool(!bool)
  }
  return (
    <nav>
      {/* Search Bar (Right) */}
      <div className='logo'>
        <span>Lite-Facebook</span>
      </div>

      {/* Logo and Icons (Center) */}
      <div className='nav-center'>
        <SearchOutlinedIcon />
        <div className='search-container'>
          <input type='search' placeholder='Search for friends, groups' />
        </div>
        <div className='icons'>
          <HomeOutlinedIcon />
          <GroupOutlinedIcon />
          <div className='notification'>
            <div onClick={handleclick}>
              <NotificationsNoneOutlinedIcon />
            </div>
            <div className='pop-out none'>{bool && <ChatPopover />}</div>
          </div>
          <MailOutlinedIcon />
        </div>
      </div>
      <img src='/images/profile-1.jpg' alt='Profile' />
    </nav>
  )
}
