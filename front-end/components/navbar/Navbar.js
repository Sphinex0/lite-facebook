'use client'
import './navbar.css'
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupAddOutlinedIcon from '@mui/icons-material/GroupAddOutlined'
import PublicOutlinedIcon from '@mui/icons-material/PublicOutlined'

export default function Navbar () {
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
          <GroupAddOutlinedIcon />
          <PublicOutlinedIcon />
        </div>
      </div>

      {/* Profile Image (End) */}
      {/* <div className='nav-image'> */}
        <img src='/images/profile-1.jpg' alt='Profile' />
      {/* </div> */}
    </nav>
  )
}
