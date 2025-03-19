'use client'
import { Group, Groups, Home, Mail, Message, Notifications } from '@mui/icons-material'
import Link from 'next/link'
import { usePathname } from 'next/navigation';
import React, { useState } from 'react'

const SideBar = () => {
    const router = usePathname();
    // const [path, setPath]= useState(router)
  
  return (
    <div className="sidebar">
        <Link href={"/"} className={`menu-item ${router == "/" && "active"}`}>
            <span><Home/></span>
            <h3>Home</h3>   
        </Link>
        <Link href={"/groups"} className={`menu-item ${router == "/groups" && "active"}`}>
            <span><Groups/></span>
            <h3>Groups</h3>
        </Link>

        <Link href={"/"} className="menu-item"  id="notifications">
            <span className='i'><Notifications/><small className="notification-count">9+</small></span>
            <h3>Notification</h3>
        </Link>

        <Link href={"/chat"} className={`menu-item ${router == "/chat" && "active"}`} id="messages-notifications">
            <span className='i'><Mail/><small className="notification-count">6</small></span>
            <h3>Messages</h3>
        </Link>

       
    </div>
  )
}

export default SideBar
