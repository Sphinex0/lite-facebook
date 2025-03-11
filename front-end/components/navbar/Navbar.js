"use client";
import './navbar.css'
import Link from "next/link";
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined';
import DarkModeOutlinedIcon from '@mui/icons-material/DarkModeOutlined';
import WbSunnyOutlinedIcon from '@mui/icons-material/WbSunnyOutlined';
import GridViewOutlinedIcon from '@mui/icons-material/GridViewOutlined';
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined';
import EmailOutlinedIcon from '@mui/icons-material/EmailOutlined';
import AccountBoxOutlinedIcon from '@mui/icons-material/AccountBoxOutlined';
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined';

export default function Navbar() {
    return (
        <div className='navbar'>
            <div className='left'>
            <Link href="/" style={{textDecoration:"none"}}>
            <span>Lite-Facebook</span>
            <HomeOutlinedIcon/>
            <DarkModeOutlinedIcon/>
            <GridViewOutlinedIcon/>
            </Link>
            </div>
            <div className='search'>
                <SearchOutlinedIcon/>
                <input type='text' placeholder='Search'/>
            </div>
            <div className='right'>
            <AccountBoxOutlinedIcon/>
            <EmailOutlinedIcon/>
            <NotificationsNoneOutlinedIcon/>
            <div className='user'>
                <img src='ergregreg.jpg' alt="no pic for now"/>
                <span>Nicolas Sad</span>
            </div>
            </div>
        </div>
    )
}