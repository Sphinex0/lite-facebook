import Link from 'next/link';
import './notification.css'
import { useEffect } from 'react';

const Notifications = ({ notifications = [], Err }) => {
  useEffect(() => {
  const Handleacceptfollow = async (user_id, Status) => {
    const response = fetch("http://localhost:8080/api/follow/decision",{
      methode: "post",
      body: JSON.stringify(
        user_id, 
        Status,)

    })
  }
  });

  return (
    <div className="notification-wrapper">
      <div className="notification-container">
        {Err && <div className="notif-err">Error loading notifications. Please try again.</div>}
        {notifications.map((notification, index) => {
      switch (notification.type) {
        case "follow-request":
          return (
            <div key={index} className="notification-div">
              <h1>Follow request</h1>
              <p>{notification.invoker} sent you a follow request</p>
              <button className="accepte" onClick={() => {
                Handleacceptfollow(notification.invoker, "accepte");
              }}>
                Accept
              </button>
              <button className="refuse" onClick={() => {
                Handleacceptfollow(notification.invoker, "refuse");
              }}>
                Reject
              </button>
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