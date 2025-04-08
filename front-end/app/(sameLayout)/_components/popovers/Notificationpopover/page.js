import Link from 'next/link';
import { useEffect, useState, useRef } from 'react';
import { FetchApi } from '@/app/helpers';
import { useRouter } from 'next/navigation';
import "./notification.css";

const Notifications = ({ notifications = [], Err }) => {
  const [items, setItems] = useState([]); 
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const containerRef = useRef();
  const redirect = useRouter();

  useEffect(() => {
    const fetchItems = async () => {
      setLoading(true);
      try {
        const res = await FetchApi(`/api/GetNotification/?page=${page}`, redirect, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
        });
        console.log("res",res)

        const newItems = await res.json();
        if (res.status === 200) {
          if (newItems.notifications != null) {
            // Filter out duplicates by checking `id`
            setItems((prev) => {
              const existingIds = new Set(prev.map(item => item.id)); 
              const uniqueNewItems = newItems.notifications.filter(
                (item) => !existingIds.has(item.id)
              );
              return [...prev, ...uniqueNewItems]; 
            });
          }
        } else {
          Err = newItems.Notifications;
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
      const container = containerRef.current;
      if (container && !loading) {
        const { scrollTop, scrollHeight, clientHeight } = container;
        if (scrollTop + clientHeight >= scrollHeight - 50) {
          setPage((prev) => prev + 1);
        }
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

  const Handlefollow = async (id, follower, status) => {
    const res = await FetchApi('/api/follow/decision', redirect, {
      method: 'POST',
      body: JSON.stringify({ follower, status }),
    });

    if (res.ok) {
      const response = await FetchApi('/api/deletenotification', redirect, {
        method: 'POST',
        body: JSON.stringify({ id }),
      });
      if (response.ok) {
        setItems((prev) => prev.filter((item) => item.id !== id));
      }
    }
  };

  const HandleInvite = async (id, invoker_id, group_id, status) => {
    const res = await FetchApi('/api/invite/decision', redirect, {
      method: 'POST',
      body: JSON.stringify({ sender: invoker_id, group_id, status }),
    });
    if (res.ok) {
      const response = await FetchApi('/api/deletenotification', redirect, {
        method: 'POST',
        body: JSON.stringify({ id }),
      });
      if (response.ok) {
        setItems((prev) => prev.filter((item) => item.id !== id));
      }
    }
  };

  return (
    <div className="notification-wrapper">
      <div className="notification-container" ref={containerRef}>
        {Err && <div className="notif-err">Error loading notifications. Please try again.</div>}
        {items && items.map((notification, index) => {
          switch (notification.type) {
            case 'follow':
              return (
                <div key={`notif${notification.id}`} className="notification-div">
                  <h1>A Follow</h1>
                  <p>You did Get a follow From a user Named {notification.invoker_name}</p>
                </div>
              );
            case 'follow-request':
              return (
                <div key={`notif${notification.id}`} className="notification-div">
                  <h1>Follow Request</h1>
                  <p><strong>{notification.invoker_name}</strong> sent you a follow request</p>
                  <button className="accepte" onClick={() => Handlefollow(notification.id, notification.invoker_id, "accepted")}>Accept</button>
                  <button className="refuse" onClick={() => Handlefollow(notification.id, notification.invoker_id, "rejected")}>Reject</button>
                </div>
              );
            case 'invitation-request':
              return (
                <div key={`notif${notification.id}`} className="notification-div">
                  <h1>Invitation Request</h1>
                  <p><strong>{notification.invoker_name}</strong> invited you to join the group <strong>{notification.group_title}</strong></p>
                  <button className="accepte" onClick={() => HandleInvite(notification.id, notification.invoker_id, notification.group_id, "accepted")}>Accept</button>
                  <button className="refuse" onClick={() => HandleInvite(notification.id, notification.invoker_id, notification.group_id, "rejected")}>Reject</button>
                </div>
              );
            case 'join':
              return (
                <div key={`notif${notification.id}`} className="notification-div">
                  <h1>Group Joining Request</h1>
                  <p><strong>{notification.invoker_name}</strong> sent you a join request to <strong>{notification.group_title}</strong></p>
                  <button className="accepte" onClick={() => HandleInvite(notification.id, notification.invoker_id, notification.group_id, "accepted")}>Accept</button>
                  <button className="refuse" onClick={() => HandleInvite(notification.id, notification.invoker_id, notification.group_id, "rejected")}>Reject</button>
                </div>
              );
            case 'event-created':
              return (
                <div key={`notif${notification.id}`} className="notification-div">
                  <Link href={`/groups/${notification.group_id}`}>
                    <h1>New Event</h1>
                  </Link>
                  <p><strong>{notification.invoker_name}</strong> created an event in <strong>{notification.group_title}</strong></p>
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