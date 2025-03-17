import React, { useEffect, useState } from 'react';
import Link from 'next/link';
import './notification.css'

const Notifications = () => {
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
  const [notificationCount, setNotificationCount] = useState(3);

  const [Err, setError] = useState("")
/*  useEffect(() => {
    const fetchNotifications = async () => {
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

    fetchNotifications();
  }, []);*/

  console.log(notificationCount);
  
  return (
    <div className="notification-wrapper">
      {notificationCount != 0 && <span className="count">{notificationCount}</span>}
      <div className="notification-container">
        {Err && <div className="notif-err">Error loading notifications. Please try again.</div>}
        {notifications.map((notification, index) => {
          switch (notification.type) {
            case "follow-request":
              return (
                <div key={index} className="notification-div">
                  <h1>Follow request</h1>
                  <p>{notification.invoker} sent you a follow request</p>
                  <button className="accepte">Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case "invitation-request":
              return (
                <div key={index} className="notification-div">
                  <h1>Invitation request</h1>
                  <p>{notification.invoker} sent you an invitation to join the group {notification.group}</p>
                  <button className="accepte">Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case "joine":
              return (
                <div key={index} className="notification-div">
                  <h1>Group joining request</h1>
                  <p>{notification.invoker} sent you a join request to {notification.group}</p>
                  <button className="accepte">Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case "event-created":
              return (
                <div key={index} className="notification-div">
                  <Link href={`/event/${notification.eventID}`}>
                    <h1>New event</h1>
                  </Link>
                  <p>{notification.invoker} created an event in {notification.group}</p>
                </div>
              );
            default:
              return null;
          }
        })}
      </div>
    </div>
  );
  
};

export default Notifications;