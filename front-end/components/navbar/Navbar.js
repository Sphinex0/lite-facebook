"use client";
import './navbar.css';
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined';
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined';
import GroupAddOutlinedIcon from '@mui/icons-material/GroupAddOutlined';
import PublicOutlinedIcon from '@mui/icons-material/PublicOutlined';

export default function Navbar() {
  return (
    <nav>
        <div className='logo'>
      <span>Lite-Facebook</span>
        </div>

      {/* Search Bar */}
      <div className="search-container">
        <SearchOutlinedIcon />
        <input type="search" placeholder="Search for friends, posts and groups" />
      </div>

      {/* Icons */}
      <div className="icons">
        <HomeOutlinedIcon />
        <GroupAddOutlinedIcon />
        <PublicOutlinedIcon />
      </div>

      {/* Profile Image */}
      <img src="/images/profile-1.jpg" alt="Profile" />
    </nav>
  );
}
