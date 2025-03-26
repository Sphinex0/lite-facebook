'use client'
import { FetchApi } from "@/app/helpers";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import style from "./group.module.css";
import { EventAvailable, EventBusy } from "@mui/icons-material";


const EventsOptions = ({ event_id }) => {
    const redirect = useRouter()
    const [going, setGoing] = useState({});
    const [notgoing, setNotGoing] = useState({});
    const handelcount = async (going) => {
        try {
            const response = await FetchApi("/api/Event/options/store", redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
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
            const response = await FetchApi("/api/Event/options/choise", redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
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
        <div className={style.options}>
            <div onClick={() => handelcount(true)} onChange={() => { }} name={`go${event_id}`} >
                <label>
                    <EventAvailable color={going.action === "action" ? "secondary" :""} />
                    going :{going.event}
                </label>
                {/* <input  /> */}
            </div>
            <div onClick={() => handelcount(false)} onChange={() => { }} name={`go${event_id}`}>
                <label>
                    <EventBusy color={notgoing.action === "action" ? "warning" :""} />
                    not going :{notgoing.event}
                </label>
                {/* <input  /> */}
            </div>
        </div>
    );
};

export default EventsOptions;