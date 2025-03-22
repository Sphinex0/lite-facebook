import Link from 'next/link';
import { useEffect, useState, useRef } from 'react';
import './notification.css';

const Notifications = ({ notifications = [], Err }) => {
  
  const [items, setItems] = useState(notifications);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
   const containerRef = useRef();

  useEffect(() => {
    const fetchItems = async () => {
      setLoading(true);
      try {
        const res = await fetch(`http://localhost:8080/api/GetNotification/?page=${page}`,{
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include', 
        });

        const newItems = await res.json();
        if (res.status == 200) {
          console.log(newItems.notifications, "new items");
          if (newItems.notification != null) {
            setItems((prev) => [...newItems.notifications, ...prev]);
          }
        } else {
          Err = newItems.Notifications
        }
      } catch (error) {
        console.error('Error fetching notifications:', error);
      }
      setLoading(false);
    };

    fetchItems();
  }, [page]);

  useEffect(() => {
    const handleScroll = () => {
      if (containerRef.current.scrollTop === 0 && !loading) {
        setPage((prev) => prev + 1);
      }
    };

    const container = containerRef.current;
    if (container) {
      container.addEventListener('scroll', handleScroll);
    }

    return () => {
      if (container) {
        container.removeEventListener('scroll', handleScroll);
      }
    };
  }, [loading]);

 const Handlefollow = async (id) => {
    
  /* if the request did get accepted or declined succesfuly then we delet it from the database*/
    const response = fetch('http://localhost:8080/api/deletenotification',{
      method: 'POST',
      body: JSON.stringify(id)
    })
  }

  return (
    <div className="notification-wrapper" ref={containerRef}>
      <div className="notification-container">
        {Err && <div className="notif-err">Error loading notifications. Please try again.</div>}
        {items && items.map((notification, index) => {
          switch (notification.type) {
            case 'follow':
              return (
                <div key={index} className="notification-div">
                  <h1>A Follow</h1>
                  <p>You did Get a follow From a user Named {notification.invoker}</p>
                  </div>
              )
            case 'follow-request':
              return (
                <div key={index} className="notification-div">
                  <h1>Follow Request</h1>
                  <p>{notification.invoker} sent you a follow request</p>
                  <button className="accepte" onClick={Handlefollow(notification.id)}>Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case 'invitation-request':
              return (
                <div key={index} className="notification-div">
                  <h1>Invitation Request</h1>
                  <p>{notification.invoker} invited you to join the group {notification.group}</p>
                  <button className="accepte">Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case 'joine':
              return (
                <div key={index} className="notification-div">
                  <h1>Group Joining Request</h1>
                  <p>{notification.invoker} sent you a join request to {notification.group}</p>
                  <button className="accepte">Accept</button>
                  <button className="refuse">Reject</button>
                </div>
              );
            case 'event-created':
              return (
                <div key={index} className="notification-div">
                  <Link href={`/event/${notification.eventID}`}>
                    <h1>New Event</h1>
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
}
export default Notifications;