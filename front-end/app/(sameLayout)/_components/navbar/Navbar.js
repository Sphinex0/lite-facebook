'use client'
import NotificationPop from '@/app/(sameLayout)/_components/popovers/Notificationpopover/page'
import Profilepop from '@/app/(sameLayout)/_components/popovers/profile/page'
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined'
import GroupOutlinedIcon from '@mui/icons-material/GroupOutlined'
import NotificationsNoneOutlinedIcon from '@mui/icons-material/NotificationsNoneOutlined'
import MailOutlinedIcon from '@mui/icons-material/MailOutlined'
import { useEffect, useState } from 'react'
import Link from 'next/link'
import { GroupsOutlined } from '@mui/icons-material'
import { usePathname, useRouter } from 'next/navigation'
import { FetchApi } from '@/app/helpers'
import "./navbar.css"
import { useWorker } from '@/app/_Context/WorkerContext'
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
export default function Navbar() {
  const [bool, setbool] = useState(false)
  const [profile, setprofile] = useState(false)
  const redirect = useRouter()
  const [notifications, setNotifications] = useState();
  const [notificationCount, setNotificationCount] = useState(0);
  const [Err, setError] = useState("")
  const { userRef } = useWorker()


  function handleclick() {
    setbool(!bool)
    setNotificationCount(0)
  }
  const handleProfileclick = (e) => {
    setprofile(!profile)
  }

  const router = usePathname();

  useEffect(() => {
    if (userRef.current.id) {
      const fetchNotifications = async () => {
        try {
          const response = await FetchApi("/api/GetNotification/?page=1", redirect, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
          });

          if (response.status == 200) {
            const data = await response.json();
            setNotifications(data.notifications || []);
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

    }
  }, [userRef.current]);

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
            {notificationCount != 0 && <span className="count">{notificationCount > 9 ? "9+" : notificationCount}</span>}
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
      <div className='notification' >
        <svg onClick={handleProfileclick} role="img" width="24px" viewBox="0 0 44 44" aria-label="icon" fill='rgb(24, 119, 242)' >
          <g xmlns="http://www.w3.org/2000/svg">
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M12.3879 6.68844C13.2206 6.68844 13.8957 6.01963 13.8957 5.19461C13.8957 4.36959 13.2206 3.70078 12.3879 3.70078C11.5552 3.70078 10.8801 4.36959 10.8801 5.19461C10.8801 6.01963 11.5552 6.68844 12.3879 6.68844ZM12.3879 8.55572C12.8641 8.55572 13.3174 8.4585 13.7287 8.283L16.3099 11.1698H5.22589V14.5312H1.83337V26.8552H5.22589V29.4692H14.2726V32.8304H1.83337V42.1668H42.1667V32.8304H29.7274V29.4692H38.7742V26.8552H42.1667V14.5312H38.7742V11.1698H27.3133L29.7998 8.38889C30.1329 8.49716 30.4887 8.55572 30.8583 8.55572C32.7319 8.55572 34.2508 7.0509 34.2508 5.19461C34.2508 3.33832 32.7319 1.8335 30.8583 1.8335C28.9847 1.8335 27.4658 3.33832 27.4658 5.19461C27.4658 6.00294 27.7538 6.74463 28.2337 7.32444L24.7954 11.1698H18.8278L15.1856 7.09632C15.5608 6.55551 15.7804 5.90047 15.7804 5.19461C15.7804 3.33832 14.2615 1.8335 12.3879 1.8335C10.5143 1.8335 8.99537 3.33832 8.99537 5.19461C8.99537 7.0509 10.5143 8.55572 12.3879 8.55572ZM36.8894 13.0371H7.11063V27.6019H36.8894V13.0371ZM14.2726 34.6977V34.6978H29.7274V34.6977H40.282V40.2996H3.71811V34.6977H14.2726ZM27.8427 32.8304V29.8429H16.1573V32.8304H27.8427ZM4.84895 16.3984H3.71811V24.9879H4.84895V16.3984ZM39.1511 16.3984H40.2819V24.9879H39.1511V16.3984ZM14.6495 20.8799C15.8986 20.8799 16.9112 19.8767 16.9112 18.6392C16.9112 17.4017 15.8986 16.3984 14.6495 16.3984C13.4004 16.3984 12.3878 17.4017 12.3878 18.6392C12.3878 19.8767 13.4004 20.8799 14.6495 20.8799ZM14.6495 22.7472C16.9395 22.7472 18.7959 20.908 18.7959 18.6392C18.7959 16.3704 16.9395 14.5312 14.6495 14.5312C12.3595 14.5312 10.5031 16.3704 10.5031 18.6392C10.5031 20.908 12.3595 22.7472 14.6495 22.7472ZM31.9892 18.6392C31.9892 19.7993 30.972 20.8799 29.539 20.8799C28.106 20.8799 27.0889 19.7993 27.0889 18.6392C27.0889 17.479 28.106 16.3984 29.539 16.3984C30.972 16.3984 31.9892 17.479 31.9892 18.6392ZM33.8739 18.6392C33.8739 20.908 31.9331 22.7472 29.539 22.7472C27.1449 22.7472 25.2041 20.908 25.2041 18.6392C25.2041 16.3704 27.1449 14.5312 29.539 14.5312C31.9331 14.5312 33.8739 16.3704 33.8739 18.6392ZM19.2667 24.0538L18.9224 23.1606L17.1617 23.8267L17.9689 25.9211H26.2366L26.8639 23.7507L25.052 23.2366L24.8158 24.0538H19.2667ZM32.3661 5.19461C32.3661 6.01963 31.691 6.68844 30.8583 6.68844C30.0256 6.68844 29.3505 6.01963 29.3505 5.19461C29.3505 4.36959 30.0256 3.70078 30.8583 3.70078C31.691 3.70078 32.3661 4.36959 32.3661 5.19461Z" fill="var(--purple)"
            >
            </path>
          </g>
        </svg>
        <div className='pop-out'>{profile && <Profilepop setprofile={setprofile} />}</div>
      </div>

    </nav>
  )
}
