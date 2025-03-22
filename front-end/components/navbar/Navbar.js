'use client'
import './navbar.css'
import NotificationPop from '@/components/popovers/Notificationpopover/page'
import Profilepop from  '@/components/popovers/profile/page'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupOutlinedIcon from '@mui/icons-material/GroupOutlined'
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined'
import MailOutlinedIcon from '@mui/icons-material/MailOutlined'
import { useEffect, useState } from 'react'
import Link from 'next/link'
import { GroupsOutlined } from '@mui/icons-material'
import { usePathname, useRouter } from 'next/navigation'
import { useWorker } from '@/app/_Context/WorkerContext'
import { FetchApi } from '@/app/helpers'
/*
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
  ]*/
export default function Navbar () {
  const [bool, setbool] = useState(false)
  const [profile, setprofile] = useState(false)
  const redirect = useRouter()
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
   
  const router = usePathname();

     useEffect(() => {
           
    
      const storedUser = JSON.parse(localStorage.getItem('user')) || {};
             setUser(storedUser);
 
    const fetchNotifications = async () => {
      try {
        const response = await FetchApi("/api/GetNotification/?page=1",redirect,{
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (response.status == 200) {
          const data = await response.json();          
          setNotifications(data.notifications);
          setNotificationCount(data.unseen); 
          
        } else {
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
    <nav className={router == "/login" || router == "/signup" ? "disable" : ""}>
      {/* Search Bar (Right) */}
      <div className='logo'>
        <Link href={'/'}>
        <span>Lite-Facebook</span>
        </Link>
      </div>

      {/* Logo and Icons (Center) */}
      <div className='nav-center '>
        <div className='icons'>
        <Link href={"/"} className={`menu-item ${router == "/" && "active"}`}>
                <span><HomeOutlinedIcon /></span>
        </Link> 
        <Link href={"/users"} className={`menu-item ${router == "/users" && "active"}`}>
                <span><GroupOutlinedIcon /></span>
            </Link>
            <Link href={"/groups"} className={`menu-item ${router == "/groups" && "active"}`}>
                <span><GroupsOutlined /></span>
            </Link>
          <div className='notification'>
            <div onClick={handleclick}>
              <NotificationsNoneOutlinedIcon />
            </div>
            {notificationCount != 0 && <span className="count">{notificationCount}</span>}
            <div className='pop-out'>{bool && <NotificationPop notifications={notifications} Err={Err} />}</div>
          
          </div>

          <Link href={"/chat"} className={`menu-item ${router == "/chat" && "active"}`} id="messages-notifications">
                <span className='i notification'>
                <MailOutlinedIcon />
                    <span className="notification-count count" id='msgCount'></span>
                </span>
            </Link>
          
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
