import { Group, Home, Mail, Message, Notifications } from '@mui/icons-material'
import React from 'react'

const SideBar = () => {
  return (
    <div className="sidebar">
        <a className="menu-item active">
            <span><Home/></span>
            <h3>Home</h3>   
        </a>
        <a className="menu-item">
            <span><Group/></span>
            <h3>Groups</h3>
        </a>

        <a className="menu-item"  id="notifications">
            <span className='i'><Notifications/><small className="notification-count">9+</small></span>
            <h3>Notification</h3>
        </a>

        <a className="menu-item" id="messages-notifications">
            <span className='i'><Mail/><small className="notification-count">6</small></span>
            <h3>Messages</h3>
        </a>

       
    </div>
  )
}

export default SideBar
