'use client'
import './navbar.css'
import NotificationPop from '@/components/popovers/Notificationpopover/page'
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupOutlinedIcon from '@mui/icons-material/GroupOutlined'
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined'
import MailOutlinedIcon from '@mui/icons-material/MailOutlined'
import { useEffect, useState } from 'react'

export default function Navbar () {
  const [bool, setbool] = useState(false)
  function handleclick () {
    setbool(!bool)
  }
  const [notifications, setNotifications] = useState([
    {
      type: 'follow-request',
      invoker: 'hamza'
    },
    {
      type: 'invitation-request',
      invoker: 'ayoub',
      group: 'programming'
    },
    {
      type: 'joine',
      invoker: 'imad',
      group: 'fitness'
    },
    {
      type: 'event-created',
      group: 'knowledge',
      invoker: 'mustafa'
    },
  ]);
  const [notificationCount, setNotificationCount] = useState(0);
  const [image, setImage] = useState("/default-profile.png")
  const [Err, setError] = useState("")

  useEffect(() => {
      const storedData = sessionStorage.getItem("Image");
      if (storedData) {
        setImage(storedData);
      }

   /* const fetchNotifications = async () => {
      try {
       
        const response = await fetch("http://localhost:8080/api/GetNotification",{
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include', 
        });

        if (response.ok) {
          const data = await response.json();
          setNotifications(data.notifications);
          setNotificationCount(data.count); 
        } else {
          setError("error while fetching notifications");
        }
      } catch (error) {
        setError("error while fetching notifications");
        console.error('Error fetching notifications:', error);
      }
    };

    fetchNotifications();*/
  }, []);

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
            {notificationCount != 0 && <span className="count">{notificationCount}</span>}
            <div className='pop-out none'>{bool && <NotificationPop notifications={notifications} Err={Err} />}</div>
          </div>
          <MailOutlinedIcon />
        </div>
      </div>
      <img  src={image} alt='Profile' />
    </nav>
  )
}
