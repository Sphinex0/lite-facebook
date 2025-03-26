
'use client';
import { FetchApi } from "@/app/helpers";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import style from "./group.module.css";
import { EventAvailable, EventBusy } from "@mui/icons-material";

const EventsOptions = ({ event_id }) => {
    const redirect = useRouter();
    const [eventData, setEventData] = useState({
        goingCount: 0,
        notGoingCount: 0,
        action: "undecided",
    });

    const fetchEventsOptions = async () => {
        try {
            const response = await FetchApi("/api/Event/options/choise", redirect, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ event_id }), 
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const data = await response.json();
            setEventData({
                goingCount: data.going_count || 0,
                notGoingCount: data.not_going_count || 0,
                action: data.action || "undecided",
            });
        } catch (error) {
            console.error("Error fetching event options:", error);
        }
    };

    const handelcount = async (going) => {
        try {
            const response = await FetchApi("/api/Event/options/store", redirect, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ event_id, going }),
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            await fetchEventsOptions();
        } catch (error) {
            console.error("Error updating event choice:", error);
        }
    };

    useEffect(() => {
        fetchEventsOptions();
    }, [event_id]);

    return (
        <div className={style.options}>
            <div onClick={() => handelcount(true)}>
                <label>
                    <EventAvailable
                        color={eventData.action === "going" ? "secondary" : "action"}
                    />
                    going: {eventData.goingCount}
                </label>
            </div>
            <div onClick={() => handelcount(false)}>
                <label>
                    <EventBusy
                        color={eventData.action === "not going" ? "warning" : "action"}
                    />
                    not going: {eventData.notGoingCount}
                </label>
            </div>
        </div>
    );
};

export default EventsOptions;