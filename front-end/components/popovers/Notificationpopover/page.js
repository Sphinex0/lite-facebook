'use client'
import './notification.css'

export default function NotificationPop () {
  const notifications = [
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
  ]
  return (
    <div className="notification-container">
      {notifications.map((notification, index) => {
        switch (notification.type) {
          case "follow-request":
            return (
              <div key={index} className="notification-div">
                <h1>Follow request</h1>
                <p>{notification.invoker} sent you a follow request</p>
                <button className='accepte'>Accept</button>
                <button className='refuse'>Reject</button>
              </div>
            );
            case "invitation-request":
            return (
              <div key={index} className="notification-div">
                <h1>Invitation request</h1>
                <p>{notification.invoker} sent you a follow request</p>
                <button className='accepte'>Accept</button>
                <button className='refuse'>Reject</button>
              </div>
            );
            case "joine":
            return (
              <div key={index} className="notification-div">
                <h1>Group joining request</h1>
                <p>{notification.invoker} sent you a join request to {notification.group}</p>
                <button className='accepte'>Accept</button>
                <button className='refuse'>Reject</button>
              </div>
            );
            case "event-created":
            return (
              <div key={index} className="notification-div">
                <h1>New event</h1>
                <p>{notification.invoker} Created an Event in {notification.group}</p>
              </div>
            );
        }
      })}
    </div>
  );
}
