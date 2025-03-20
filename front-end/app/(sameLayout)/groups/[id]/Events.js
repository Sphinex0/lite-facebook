'use client'

import { useState, useEffect } from "react";

const Events = ({ groupID }) => {
    const [eventsData, setEventsData] = useState([]); 
    const [loading, setLoading] = useState(true); 
    const [error, setError] = useState(null); 

    useEffect(() => {
      
        const fetchEvents = async () => {
            try {
                const response = await fetch("http://localhost:8080/api/Events", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    credentials: "include",
                    body: JSON.stringify({ group_id: parseInt(groupID) }),
                });

                console.log(JSON.stringify({ group_id: parseInt(groupID) }));

                if (response.ok) {
                    const data = await response.json();
                    console.log(data);
                    setEventsData(data);
                    setLoading(false);  
                } else {
                    setError("Failed to fetch events");
                    setLoading(false);
                }
            } catch (error) {
                setError("An error occurred while fetching events.");
                setLoading(false);
            }
        };

        fetchEvents(); 
    }, [groupID]);

    if (loading) {
        return <div>Loading...</div>;
    }
    if (error) {
        return <div>{error}</div>; 
    }

    return (
        <div>
            {eventsData.length === 0 ? (
                <div>No events found.</div> 
            ) : (
                <ul>
                    {eventsData.map((event, index) => (
                    <div className='feed'>
                    <div className="user">
                        <div className="profile-photo">
                            <img src="./images/profile-13.jpg" />
                        </div>
                        <div className="ingo">
                            <h3>{event.title} </h3>  
                            <small>{event.description}</small> 
                        </div>
                    </div>
                  </div>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default Events;
