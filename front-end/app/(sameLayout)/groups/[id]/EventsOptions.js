'use client'
import { useEffect, useState } from "react";


const EventsOptions = ({ event_id }) => {
    const [going, setGoing] = useState({});
    const [notgoing, setNotGoing] = useState({});
    const handelcount = async (going) => {
      
        try {
            const response = await fetch("http://localhost:8080/api/Event/options/store", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include",
                body: JSON.stringify({ event_id, going: going })
            });
            console.log(response);
            
            if (response.ok) {
                fetchEventsOptions(true);
                fetchEventsOptions(false);
            }
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
        } catch (error) {
            throw new Error('Network response was not ok');
        }

    }


    const fetchEventsOptions = async (going) => {
        try {
            const response = await fetch("http://localhost:8080/api/Event/options/choise", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include",
                body: JSON.stringify({ event_id, going: going })
            });
         

            if (!response.ok) {
                throw new Error('Network response was not ok');
            } else {
                const data = await response.json();
                console.log(data);
                console.log(going);
                if (data !== null) {
                    if (going) {
                        setGoing(data);
                    } else {
                        setNotGoing(data);
                    }
                }
            }
        } catch (error) {
            throw new Error('Network response was not ok');
        }
    }

    useEffect(() => {
        fetchEventsOptions(true);
        fetchEventsOptions(false);
    }, [event_id]);
    return (
        <div>
            <label>going :{going.event}</label>
            <input onClick={() => handelcount(true)} onChange={()=>{}}  checked={going.action === "action"  && true} type="radio" name={`go${event_id}`} /><br/>
            <label>not going :{notgoing.event}</label>
            <input onClick={() => handelcount(false)} onChange={()=>{}}  checked={notgoing.action === "action" && true}  type="radio" name={`go${event_id}`}  />
        </div>
    );
};

export default EventsOptions;