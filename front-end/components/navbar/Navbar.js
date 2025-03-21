'use client'
import './navbar.css'
import NotificationPop from '@/components/popovers/Notificationpopover/page'
import Profilepop from  '@/components/popovers/profile/page'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupOutlinedIcon from '@mui/icons-material/GroupOutlined'
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined'
import MailOutlinedIcon from '@mui/icons-material/MailOutlined'
import { useEffect, useState } from 'react'
/*[
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
  ]*/
export default function Navbar () {
  const [bool, setbool] = useState(false)
  const [profile, setprofile] = useState(false)
  function handleclick () {
    setbool(!bool)
  }
  function handleProfileclick () {
    setprofile(!profile)
  }
  const [notifications, setNotifications] = useState();
  const [notificationCount, setNotificationCount] = useState(0);
  const [Err, setError] = useState("")
   const [user, setUser] = useState({});
 
     useEffect(() => {
             const storedUser = JSON.parse(localStorage.getItem('user')) || {};
             setUser(storedUser);
 
    const fetchNotifications = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/GetNotification/1",{
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include', 
        });

        if (response.status == 200) {
          const data = await response.json();          
          setNotifications(data.notifications);
          setNotificationCount(data.count); 
        } else {
          console.log(response,"waaaaaaaaaaaaaaaaaaaakwaaaaaaaaaaaaaaaaaak");
          
          setError("error while fetching notifications");
        }
      } catch (error) {
        setError("error while fetching notifications");
        console.error('Error fetching notifications:', error);
      }
    };

    fetchNotifications();
  }, []);

  return (
    <nav>
      {/* Search Bar (Right) */}
      <div className='logo'>
        <span>Lite-Facebook</span>
      </div>

      {/* Logo and Icons (Center) */}
      <div className='nav-center'>
        <div className='icons'>
          <HomeOutlinedIcon />
          <GroupOutlinedIcon />
          <div className='notification'>
            <div onClick={handleclick}>
              <NotificationsNoneOutlinedIcon />
            </div>
            {notificationCount != 0 && <span className="count">{notificationCount}</span>}
            <div className='pop-out'>{bool && <NotificationPop notifications={notifications} Err={Err} />}</div>
          </div>
          <MailOutlinedIcon />
        </div>
      </div>
      <div className='notification'>
      <div onClick={handleProfileclick}>
      <img src={user.image || "/default-profile.png"} alt='Profile' />
      </div>
      <div className='pop-out'>{profile && <Profilepop/>}</div>
      </div>
    </nav>
  )
}
