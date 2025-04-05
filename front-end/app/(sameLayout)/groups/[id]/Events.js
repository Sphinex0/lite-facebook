'use client'
import { FetchApi } from "@/app/helpers";
import style from "./group.module.css";
import { Add, DisabledByDefault, EventAvailable } from "@mui/icons-material";
import EventsOptions from "./EventsOptions";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

const Events = ({ groupID }) => {
    const [eventsData, setEventsData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [title, setTitle] = useState('');
    const [going, setGoing] = useState(false)
    const [description, setDescription] = useState('');
    const [day, setDay] = useState('');
    const redirect = useRouter()

    const CreateGroup = () => {
        const element = document.querySelector('#formId');
        element.style.display = "flex";
    }

    const SeeClick = () => {
        const element = document.querySelector('#formId');
        element.style.display = "none";
    }

    const handleSubmit = async (e) => {
        e.preventDefault();

        setTitle("");
        setDescription("");

        try {
            const response = await FetchApi("/api/Event/store", redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    event: { group_id: parseInt(groupID), Title: title, Description: description, day },
                    going
                })
            });

            if (!response.ok) {
                setError("An error occurred while submitting the event.");
                setLoading(false);
            } else {
                fetchEvents()
                SeeClick()
            }
        } catch (error) {
            setError("An error occurred while submitting the event.");
            setLoading(false);
        }
    };

    const fetchEvents = async () => {
        try {
            const response = await FetchApi("/api/Events", redirect, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ group_id: parseInt(groupID) }),
            });


            if (response.ok) {
                const data = await response.json();
                console.log(data)
                setEventsData(data);
                setLoading(false);
            } else {
                // setError("Failed to fetch events");
                setError(".");
                setLoading(false);
            }
        } catch (error) {
            setError("An error occurred while fetching events.");
            setLoading(false);
        }
    };

    useEffect(() => {
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
            <div>
                <button className={style.create} onClick={CreateGroup}>
                    <Add />
                </button>
            </div>

            <div className={style.formclass} id="formId">
                <div onClick={SeeClick}>
                    <DisabledByDefault />
                </div>
                <form method='POST' className={style.formAddGroup} aria-multiselectable onSubmit={handleSubmit}>
                    <label htmlFor='title'>Title</label>
                    <input type='text' className={style.InputTitle} id='title' value={title} onChange={(e) => setTitle(e.target.value)} />
                    <label htmlFor='description'>Description</label>
                    <input type='text' className={style.InputDescriptopn} id='description' value={description} onChange={(e) => setDescription(e.target.value)} />
                    <input type='datetime-local' className={style.InputDescriptopn} id='day' value={day} onChange={(e) => setDay(e.target.value)} />
                    
                    <div className={style.options}>
                        <div>
                            <label onClick={() => setGoing(true)} >
                                <EventAvailable
                                    color={going ? "secondary" : "action"}
                                />
                                going
                            </label>
                        </div>
                        <div>
                            <label onClick={() => setGoing(false)} >
                                <EventAvailable
                                    color={going == false ? "warning" : "action"}
                                />
                                Not going
                            </label>
                        </div>
                    </div>

                    <button className={style.InputButton} type='submit'>Submit</button>
                </form>
            </div>

            <div>
                {eventsData === null ? (
                    <div>No events found.</div>
                ) : (
                    <div className={style.events}>
                        {eventsData.map((event, index) => (
                            <div className={style.event} key={index}>
                                <div className={style.user}>
                                    <div className={style.info}>
                                        <h3>{event.title}</h3>
                                        <small>{event.description}</small>

                                        <div>
                                            <EventsOptions event_id={event.id} />
                                        </div>
                                        <h6>{event.day}</h6>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
};

export default Events;
